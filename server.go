package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"strings"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
)

type server struct {
	rooms    map[string]*room
	commands chan command
}

func newServer() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command),
	}
}

func (s *server) run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_NICK:
			s.nick(cmd.client, cmd.args)
		case CMD_JOIN:
			s.join(cmd.client, cmd.args)
		case CMD_ROOMS:
			s.listRooms(cmd.client)
		case CMD_MSG:
			s.msg(cmd.client, cmd.args)
		case CMD_QUIT:
			s.quit(cmd.client)
		}
	}
}

func (s *server) newClient(conn net.Conn) {
	log.Printf("new client connected: %s", conn.RemoteAddr().String())

	c := &client{
		conn:     conn,
		nick:     fmt.Sprintf("anon-%s", s.RandomName(10)),
		commands: s.commands,
	}

	c.readInput()
}

func (s *server) RandomName(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func (s *server) nick(c *client, args []string) {
	if len(args) < 2 {
		c.msg("nick name is required. Usage: /nick NAME")
		return
	}
	c.nick = strings.TrimSpace(args[1])
	c.msg(fmt.Sprintf("Your name is set to %s", c.nick))
}

func (s *server) join(c *client, args []string) {
	if len(args) < 2 {
		c.msg("room name is required. Usage: /join ROOM_NAME")
		return
	}

	roomName := strings.TrimSpace(args[1])
	if !strings.HasPrefix(roomName, "#") {
		c.msg(fmt.Sprintf("room name must begin with #. e.g: /join #%s", roomName))
		return
	}

	r, ok := s.rooms[roomName] // check if room already exist
	if !ok {                   // create room if it doesn't exist
		r = &room{
			name:    roomName,
			members: make(map[net.Addr]*client),
		}
		s.rooms[roomName] = r
	}

	if c.room != nil && c.room.name == r.name { // if user is already in the room
		c.msg(fmt.Sprintf("you are already in room %s", r.name))
		return
	}

	s.quitCurrentRoom(c) // remove user from previous room

	r.members[c.conn.RemoteAddr()] = c // add member
	c.room = r

	r.broadcast(c, fmt.Sprintf("%s has joined the room", c.nick))
	c.msg(fmt.Sprintf("welcome to %s", r.name))
}

func (s *server) listRooms(c *client) {
	var rooms []string
	for name := range s.rooms {
		rooms = append(rooms, name)
	}
	if len(rooms) == 0 {
		c.msg("no available rooms")
		return
	}
	c.msg(fmt.Sprintf("available rooms are: %s", strings.Join(rooms, ", ")))
}

func (s *server) msg(c *client, args []string) {
	if c.room == nil {
		c.err(fmt.Errorf("you must join the room first"))
		return
	}
	message := strings.Join(args[1:], " ")
	c.room.broadcast(c, fmt.Sprintf("%s: %s", c.nick, message))
}

func (s *server) quit(c *client) {
	log.Printf("client has disconnected: %s", c.conn.RemoteAddr().String())
	s.quitCurrentRoom(c)
	c.msg("until next time!")
	c.conn.Close()
}

func (s *server) quitCurrentRoom(c *client) {
	if c.room != nil {
		delete(c.room.members, c.conn.RemoteAddr())
		c.room.broadcast(c, fmt.Sprintf("%s has left the room", c.nick)) // broadcast to all room members
		c.room = nil
	}
}
