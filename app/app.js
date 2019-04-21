const express = require('express')
const CronJob = require('cron').CronJob;
const ArticlesModel = require('./models/articles_model')

const app = express()

new CronJob('0 * * * *', async () => {
  try {
    const articlesModel = new ArticlesModel()
    const results = await articlesModel.add()
    console.log('success to add new articles')
    console.log(results)
  } catch (error) {
    console.log(error.message)
  }
}, null, true, 'Asia/Tokyo');

module.exports = app
