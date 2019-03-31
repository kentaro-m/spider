var express = require('express');
var router = express.Router();
const ArticlesModel = require('../models/articles_model')

/* GET home page. */
router.get('/add', async function(req, res, next) {
  const articlesModel = new ArticlesModel()
  await articlesModel.add()
  res.send('respond with a resource');
});

module.exports = router;
