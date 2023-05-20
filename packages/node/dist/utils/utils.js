"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.mapToObject = void 0;
var mapToObject = function (input) {
    var content = input.match(/\[(.*?)\]/)[1];
    var pairs = content.split(" ").map(function (pair) { return pair.split(":"); });
    var output = pairs.reduce(function (obj, _a) {
        var key = _a[0], value = _a[1];
        //@ts-ignore
        obj[key] = value;
        return obj;
    }, {});
    console.log(output);
};
exports.mapToObject = mapToObject;
