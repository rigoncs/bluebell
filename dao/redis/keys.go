package redis

// redis key 注意使用命名空间的方式，方便查询和拆分
const (
	KeyPrefix          = "bluebell:"
	KeyPostTimeZSet    = "post:time"
	KeyPostScoreZSet   = "post:score"
	KeyPostVotedZSetPF = "post:voted:" //zset；记录用户及投票类型，参数是post_id

	KeyCommunitySetPF = "community:" //set;保存每个分区下帖子的id
)
