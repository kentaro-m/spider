const uuidv4 = require('uuid/v4')
const moment = require('moment-timezone')
const ArticleEntity = require('../entities/article_entity')
const { getArticles, postArticle, filterNewArticles } = require('../lib/utils')
const config = require('config')

module.exports = class ArticlesModel {
  constructor () {
  }

  async add () {
    try {
      const articles = await getArticles(config.sources)

      const newArticles = filterNewArticles(articles, 1)

      if (newArticles.length === 0) {
        throw new Error('not exist new articles')
      }

      for (const article of newArticles) {
        const articleEntity = new ArticleEntity(
          uuidv4(),
          article.title,
          article.link,
          moment(article.pubDate).tz("Asia/Tokyo"),
          moment().tz("Asia/Tokyo"),
          moment().tz("Asia/Tokyo")
        )

        await postArticle(articleEntity)
      }
    } catch (error) {
      throw error
    }
  }
}
