const Parser = require('rss-parser')
const axios = require('axios')
const moment = require('moment')

require('dotenv').config()

async function fetchArticlesByBlog (sources) {
  const parser = new Parser()

  const articles = await Promise.all(sources.map(async (source) => {
    const result = await axios.get(source.url)
    const feed = await parser.parseString(result.data)
    return feed
  }))

  return articles
}

async function getArticles (sources) {
  const blogsList = await fetchArticlesByBlog(sources)

  const articles = []

  blogsList.map(articlesList => {
    articlesList.items.map(item => {
      articles.push(item)
    })
  })

  return articles
}

async function postArticle (article) {
  const options = {
    url: '/articles',
    method: 'post',
    baseURL: `${process.env.API_PROTOCOL}://${process.env.API_HOST}:${process.env.API_PORT}/`,
    data: {
      title: article.title,
      url: article.url,
      pub_date: article.pubDate,
    }
  }

  const response = await axios(options)

  return response
}

function filterNewArticles (articles, since) {
  const newArticles = articles.filter(article => {
    return moment(article.pubDate)
        .tz("Asia/Tokyo")
        .isAfter(
            moment()
            .subtract(since, 'hours')
            .tz("Asia/Tokyo"),
            "minutes"
        )
  })
  return newArticles
}

module.exports = {
  getArticles,
  postArticle,
  filterNewArticles
}
