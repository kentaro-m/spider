const uuidv4 = require('uuid/v4');
const moment = require('moment');
const ArticlesRepositry = require('../repositories/articles_repository')
const ArticleEntity = require('../entities/article_entity')
const { getArticles } = require('../lib/utils')
const config = require('config')

module.exports = class ArticlesModel {

  constructor() {
  }

  async add() {
    try {
      const articlesRepositry = new ArticlesRepositry()

      const articles = await getArticles(config.feedUrls)

      for (const article of articles) {
        const articleEntity = new ArticleEntity(
          uuidv4(),
          article.title,
          article.link,
          moment(article.pubDate).format('YYYY-MM-DD HH:mm:ss')
        )

        await articlesRepositry.create(articleEntity)
      }

      articlesRepositry.end()
    } catch (error) {
      throw error
    }
  }

  async get() {

  }
}