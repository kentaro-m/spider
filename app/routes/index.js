const express = require('express')
const router = express.Router()
const ArticlesModel = require('../models/articles_model')

/* GET home page. */
router.get('/crawl', async function (req, res, next) {
  try {
    const articlesModel = new ArticlesModel()
    const results = await articlesModel.crawl(req)
    console.log(results)
    res.send('success to add new articles')
  } catch (error) {
    console.log(error)
    res.send(error.message)
  }
})

module.exports = router
