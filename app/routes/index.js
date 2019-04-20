const express = require('express')
const router = express.Router()
const ArticlesModel = require('../models/articles_model')

/* GET home page. */
router.get('/add', async function (req, res, next) {
  try {
    const articlesModel = new ArticlesModel()
    await articlesModel.add()
    res.send('success to add new articles')
  } catch (error) {
    res.send(error.message)
  }
})

module.exports = router
