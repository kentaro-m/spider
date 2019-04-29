const Parser = require('rss-parser')
const axios = require('axios')
const moment = require('moment-timezone')

module.exports = class ArticlesCrawler {
  constructor(options) {
    this.options = options
    this.items = []
  }

  async _fetch() {
    const parser = new Parser()

    const results = await Promise.all(this.options.sources.map(async (source) => {
      const result = await axios.get(source.url)
      const feed = await parser.parseString(result.data)
      return feed
    }))

    results.map(data => {
      data.items.map(item => {
        this.items.push(item)
      })
    })
  }

  async _filter() {
    this.items = this.items.filter(item => {
      return moment(item.pubDate)
        .tz("Asia/Tokyo")
        .isAfter(
          moment()
            .subtract(this.options.since, 'hours')
            .tz("Asia/Tokyo"),
          "minutes"
        )
    })
  }

  async get() {
    await this._fetch()
    await this._filter()
    return this.items
  }
}