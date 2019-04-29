const moment = require('moment-timezone')
const ArticlesCrawler = require('../lib/articles_crawler')
const ArticlesClient = require('../lib/articles_client')
const config = require('config')

require('dotenv').config()

module.exports = class ArticlesModel {
  constructor () {
  }

  async crawl(req) {
    try {
      const articlesCrawler = new ArticlesCrawler({
        sources: config.crawler.sources,
        since: req.query.since || 1
      })
      const articles = await articlesCrawler.get()

      if (articles.length === 0) {
        throw new Error('not exist new articles')
      }

      let baseUrl = `${process.env.API_PROTOCOL}://${process.env.API_HOST}/`

      if (process.env.API_PORT) {
        baseUrl = `${process.env.API_PROTOCOL}://${process.env.API_HOST}:${process.env.API_PORT}/`
      }

      const articlesClient = new ArticlesClient({
        baseUrl
      })
      const results = []

      for (const article of articles) {
        const data = {
          title: article.title,
          url: article.link,
          pub_date: moment(article.pubDate).tz("Asia/Tokyo"),
        }

        results.push(data)
        await articlesClient.post(data)
      }

      return results
    } catch (error) {
      throw error
    }
  }
}
