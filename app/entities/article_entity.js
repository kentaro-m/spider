module.exports = class ArticleEntity {

  constructor (id, title, url, pubDate, createdAt, updatedAt) {
    this._id = id
    this._title = title
    this._url = url
    this._pubDate = pubDate
    this._createdAt = createdAt
    this._updatedAt = updatedAt
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

  get createdAt() {
    return this._createdAt;
  }

  get updatedAt() {
    return this._updatedAt;
  }
}
