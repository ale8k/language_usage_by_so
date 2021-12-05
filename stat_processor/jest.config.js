const tsJest = require("ts-jest/jest-preset");

module.exports =  {
  ...tsJest,
  globals: {
    name: "Diglett man",
    testEnvironment: "node",
    reporters: ["default"],
  }
};