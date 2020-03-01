'use strict'
const excel = require('convert-excel-to-json');
const fs = require('fs');
const path = require('path');
const parseOrd = require('./parseOrder');
const express = require('express');
const app = express();
const multer = require('multer');
//enums
const InputTypeEnum = require('./InputTypeEnum');
const AnswerTypeEnum = require('./AnswerTypeEnum');

app.use(express.static(__dirname));
app.use(multer({dest:"./"}).single("file"));
app.post('/', (req, res, next) => {
  let filedata = req.file;
  let namefile = filedata.filename;
  const result = excel({
    sourceFile: "./"+namefile
  });

  let parsOr = new parseOrd(result["Orders"]);
  let jsonParse = parsOr.toConfig();
  console.log(jsonParse);
  fs.unlinkSync("./"+namefile);
  res.send(parsOr.fileterJson(jsonParse)); 
});

app.get("/", (req, res) => {
  res.send("its works");
});

app.listen(3000, "10.55.124.223");


