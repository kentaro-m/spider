const Parser = require('rss-parser')
const axios = require('axios')

async function fetchArticlesByBlog (urls) {
  const parser = new Parser()

  const articles = await Promise.all(urls.map(async (url) => {
    const result = await axios.get(url)
    const feed = await parser.parseString(result.data)
    return feed
  }))

  return articles
}

async function getArticles (urls) {
  const blogsList = await fetchArticlesByBlog(urls)

  const articles = []

  blogsList.map(articlesList => {
    articlesList.items.map(item => {
      articles.push(item)
    })
  })

  return articles
}

module.exports = {
  getArticles
}
