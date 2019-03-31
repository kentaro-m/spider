const BaseRepository = require('./base_repository')

module.exports = class ArticlesRepository extends BaseRepository {
  constructor () {
    super()
  }

  async create (article) {
    const sql = `
      INSERT INTO
        articles
      SET
        id = ?,
        title = ?,
        url = ?,
        pub_date = ?;
    `

    const result = await this.connection.query(sql, [
      article.id,
      article.title,
      article.url,
      article.pubDate
    ])

    return result
  }
}
