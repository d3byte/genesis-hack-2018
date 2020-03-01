let names = require('./nameCells');
module.exports = class parseOrd {
  parseObjEx;
  questions;
  objExcel;
  constructor(objExcel) {
    this.objExcel = objExcel;
    this.parseObjEx = this.parseOrder(objExcel);
    this.questions = this.createQuestions();
  }

   parseOrder (orders) {
    let res = new Object();
    let lenSizeVal = 0;
    for (var key in orders[1]) res[key] = [];
    for (var key in orders) lenSizeVal++;

    for (var key in res) {
      for (var i = 0; i <= lenSizeVal; i++) {
        for (var keyj in orders[i]) {
          if (keyj == key && orders[i][keyj] != "Таблица с данными" && i != 1) {
            res[key].push(orders[i][keyj]);
          }
        }
      }
    }
    return res;
  }

  getAnswers() {
    let answer = new Array();
    let answerOut = new Array();
    let tmp = [];

    for (var i = 2; i < this.objExcel.length; i++) {
        for (var key in this.objExcel[i]) {
          if (key != "B" && key != "C" && key != "A") {
            tmp.push(this.objExcel[i][key]);
            
        }
    }
    answer.push(tmp);
    tmp = [];
  }
  let answerOtmp = [];
  for (var i = 0; i < answer.length; i++) {
      for (var j = 0; j < answer[i].length; j+=2) {
        answerOtmp.push({
          text: answer[i][j],
          value: answer[i][j+1]
        })
      }
      answerOut.push(answerOtmp);
      answerOtmp = [];
  }
  //console.log(answerOut);
  return answerOut;
}

  createQuestions() {
    let parser = this.parseObjEx;
    let questions = new Array();
    let getAnswers = this.getAnswers();
    const newParser = {};
    for (var key in names) {
      for (var keyj in parser) {
        if (key == keyj) {
          delete Object.assign(newParser, parser, {[names[key]]: parser[keyj] })[keyj];
        }
      }
      delete parser[key];
    }

    let sizeq = 0;
    for (var key in newParser["question"]) {
      sizeq++
    }
   
    for (var i = 0; i < sizeq; i++) {
      questions.push({
        question: newParser["question"][i],
        answerType: newParser["answerType"][i],
        answers: getAnswers[i]
      });
    }
   
    for (var key in questions) {
      if (questions[key]["answerType"] == "Rate") {
          questions[key]["options"] = {
            rateScales: ['Плохо', 'Отлично']
          }
      }
    }
    this.questions = questions;
    return questions;
  }

  toConfig() {
      return {
        title: "",
        publicToken: "",
        questions: this.questions
      };
  }

  //rules view json
  fileterJson(json) {
    for (var i in json["questions"]) 
      if (json["questions"][i]["answerType"] == "Input") json["questions"][i]["answers"] = [];

      return json;
  }
 
}

