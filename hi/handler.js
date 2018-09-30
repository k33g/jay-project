"use strict"

let message = {
    who: "Jay"
  , what: "🤖"
  , text: "👋 I 😍 OpenFaaS"
}

module.exports = (context, callback) => {
  callback(undefined, message);
}
