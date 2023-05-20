"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var Elix = /** @class */ (function () {
    function Elix(name) {
        this.name = "user";
        if (name) {
            this.name = name;
        }
    }
    Elix.prototype.greet = function () {
        return "hello ".concat(this.name, "!");
    };
    return Elix;
}());
exports.default = Elix;
