package util

const (
	GET_MOVIE = "select * from m_movie"
	GET_MOVIE_BY_ID = "select * from m_movie where movie_id = ?"
	POST_MOVIE = "insert into m_movie values(?,?,?,?,?)"
	UPDATE_MOVIE = `update m_movie set movie_title=?,movie_duration=?,movie_imageUrl=?,movie_synopsis=? where movie_id=?;`
	DELETE_MOVIE = `delete from m_movie where movie_id=?`
)
