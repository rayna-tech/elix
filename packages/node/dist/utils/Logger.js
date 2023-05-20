"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var chalk_1 = require("chalk");
exports.default = (function (msg) {
    var p = " Elix ";
    return {
        info: function () {
            console.log("".concat(chalk_1.default.bgCyanBright(p), " ").concat(msg));
        },
        error: function () {
            console.log("".concat(chalk_1.default.bgRedBright(p), " ").concat(msg));
        },
        rest: function () {
            console.log("".concat(chalk_1.default.bgMagentaBright(p), " ").concat(msg));
        },
        warn: function () {
            console.log("".concat(chalk_1.default.bgYellowBright(p), " ").concat(msg));
        }
    };
});
