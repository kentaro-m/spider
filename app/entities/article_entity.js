const BaseEntity = require('./base_entity')

module.exports = class ArticleEntity extends BaseEntity {
  constructor (id, title, url, pubDate) {
    super()
    this._id = id
    this._title = title
    this._url = url
    this._pubDate = pubDate
  }

  get id () {
    return this._id
  }

  get title () {
    return this._title
  }

  get url () {
    return this._url
  }

  get pubDate () {
    return this._pubDate
  }
}
