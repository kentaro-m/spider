const config = require('config')
const mysql = require('mysql')

module.exports = class BaseRepository {
  constructor () {
    this.connection = mysql.createConnection({
      host: config.db.host,
      user: config.db.user,
      password: config.db.password,
      database: config.db.database,
      port: config.db.port,
      charset: config.db.charset
    })
  }

  end () {
    this.connection.end()
  }

  async create () {
    throw new Error('Method not implemented.')
  }

  async update () {
    throw new Error('Method not implemented.')
  }

  async delete () {
    throw new Error('Method not implemented.')
  }

  async find () {
    throw new Error('Method not implemented.')
  }

  async findOne () {
    throw new Error('Method not implemented.')
  }
}
