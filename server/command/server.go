//
// server/command/server.go
//
// Copyright (c) 2017 Junpei Kawamoto
//
// This file is part of cgss.
//
// cgss is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// cgss is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with cgss.  If not, see <http://www.gnu.org/licenses/>.
//

package command

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"

	"github.com/itslab-kyushu/cgss/kvs"
)

// Server defines a KVS server.
type Server struct {
	// Root is a path to the document root.
	Root string
}

// Get returns a value associated with the given key.
func (s *Server) Get(ctx context.Context, key *kvs.Key) (res *kvs.Value, err error) {

	target := filepath.Join(s.Root, filepath.ToSlash(key.Name))
	info, err := os.Stat(target)
	if err != nil {
		return
	} else if info.IsDir() {
		return nil, fmt.Errorf("The given key is a bucket name")
	}

	data, err := ioutil.ReadFile(target)
	if err != nil {
		return
	}

	res = &kvs.Value{}
	if err = proto.Unmarshal(data, res); err != nil {
		return
	}

	return res, nil
}

// Put stores a given entry as a file.
func (s *Server) Put(ctx context.Context, entry *kvs.Entry) (*kvs.PutResponse, error) {

	target := filepath.Join(s.Root, filepath.ToSlash(entry.Key.Name))
	info, err := os.Stat(target)
	if err == nil && info.IsDir() {
		return nil, fmt.Errorf("The given key is used as a bucket name")
	}

	data, err := proto.Marshal(entry.Value)
	if err != nil {
		return nil, err
	}
	return &kvs.PutResponse{}, ioutil.WriteFile(target, data, 0644)

}

// Delete deletes a given file.
func (s *Server) Delete(ctx context.Context, key *kvs.Key) (*kvs.DeleteResponse, error) {

	target := filepath.Join(s.Root, filepath.ToSlash(key.Name))
	info, err := os.Stat(target)
	if err != nil {
		return nil, err
	} else if info.IsDir() {
		return nil, fmt.Errorf("Given key is not associated with any items")
	}

	// TODO: Delete empty directories.
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		return &kvs.DeleteResponse{}, os.Remove(target)
	}

}

// List lists up items stored in this KVS.
func (s *Server) List(_ *kvs.ListRequest, server kvs.Kvs_ListServer) error {

	return filepath.Walk(s.Root, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		item, err := filepath.Rel(s.Root, path)
		if err != nil {
			return err
		}
		return server.Send(&kvs.Key{
			Name: item,
		})

	})

}
