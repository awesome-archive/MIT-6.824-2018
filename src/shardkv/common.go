package shardkv

import "shardmaster"
import "strconv"
//
// Sharded key/value server.
// Lots of replica groups, each running op-at-a-time paxos.
// Shardmaster decides which group serves each shard.
// Shardmaster may change shard assignment from time to time.
//
// You will have to modify these definitions.
//

const (
	OK            = "OK"
	ErrNoKey      = "ErrNoKey"
	ErrWrongGroup = "ErrWrongGroup"
)

type Err string

// Put or Append
type PutAppendArgs struct {
	Me    int64
	MsgId int64
	Key   string
	Value string
	Op    string // "Put" or "Append"
	Shard int
}

type PutAppendReply struct {
	WrongLeader bool
	Err         Err
}

type GetArgs struct {
	Key string
	Shard int
}

type GetReply struct {
	WrongLeader bool
	Err         Err
	Value       string
}

type ReqShared struct {
	Shards []int
	Config shardmaster.Config
}

type RespShared struct {
	Successed bool
	Config  shardmaster.Config
	Data    map[int]map[string]string
	MsgIDs  map[int64] int64
}

type RespShareds struct {
	ConfigNum  int
	Shards     map[int]RespShared
}

type ReqDeleteShared struct {
	Shard int
	Config shardmaster.Config
}

type RespDeleteShared struct {
	Shard int
	Config shardmaster.Config
}

type GroupShards struct {
	Shards map[int][]int
	Config shardmaster.Config
}

func GetGroupShards(Shards *[shardmaster.NShards]int, group int) map[int]int {
	rst := make(map[int]int)
	for i := 0; i < len(*Shards); i++ {
		if (*Shards)[i] == group {
			rst[i] = group
		}
	}
	return rst
}

func GetGroupShardsString(shards map[int][]int) (rst string ){
	for key,value := range shards {
		rst += strconv.Itoa(key)
		rst += "{"
		for i:=0;i<len(value);i++ {
			rst += strconv.Itoa(value[i])
			rst += ","
		}
		rst += "}"
	}
	return rst
}