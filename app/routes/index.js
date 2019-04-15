const express = require('express')
const router = express.Router()
const ArticlesModel = require('../models/articles_model')

/* GET home page. */
router.get('/add', async function (req, res, next) {
  try {
    const articlesModel = new ArticlesModel()
    const response = await articlesModel.add()
    res.send(response)
  } catch (error) {
    res.send(error.message)
  }
})

module.exports = router
