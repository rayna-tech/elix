"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;
    return g = { next: verb(0), "throw": verb(1), "return": verb(2) }, typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (g && (g = 0, op[0] && (_ = 0)), _) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
Object.defineProperty(exports, "__esModule", { value: true });
var axios_1 = require("axios");
var Logger_1 = require("../utils/Logger");
var Client = /** @class */ (function () {
    function Client(options) {
        options.url = options.url.endsWith("/")
            ? options.url
                .split("/")
                .slice(0, options.url.split("/").length - 1)
                .join("/")
            : options.url;
        this.options = options;
        if (this.options.seedData) {
            this.seed();
        }
    }
    Client.prototype.set = function (key, value) {
        return __awaiter(this, void 0, void 0, function () {
            var req;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0: return [4 /*yield*/, axios_1.default
                            .post("".concat(this.options.url, "/new"), JSON.stringify({
                            key: key,
                            store: value,
                        }))
                            .then(function (r) {
                            return {
                                ok: true,
                                data: r.data,
                            };
                        })
                            .catch(function (r) {
                            return {
                                ok: false,
                                reason: "".concat(r),
                            };
                        })];
                    case 1:
                        req = _a.sent();
                        if (!req.ok) {
                            //@ts-ignore
                            (0, Logger_1.default)("Error setting key value. Please check the db server terminal for more information.").error();
                        }
                        else {
                            (0, Logger_1.default)("Key: ".concat(key, " has been assigned to the value: ").concat(JSON.stringify(value)));
                        }
                        return [2 /*return*/];
                }
            });
        });
    };
    Client.prototype.get = function (key) {
        return __awaiter(this, void 0, void 0, function () {
            var req;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0: return [4 /*yield*/, axios_1.default
                            .get("".concat(this.options.url, "/get/").concat(key))
                            .then(function (r) {
                            return {
                                ok: true,
                                data: r.data,
                            };
                        })
                            .catch(function (r) {
                            return {
                                ok: false,
                                reason: "".concat(r),
                            };
                        })];
                    case 1:
                        req = _a.sent();
                        if (!req.ok) {
                            //@ts-ignore
                            (0, Logger_1.default)("Error getting key value. Please check the db server terminal for more information.").error();
                        }
                        else {
                            //@ts-ignore
                            return [2 /*return*/, req.data];
                        }
                        return [2 /*return*/];
                }
            });
        });
    };
    Client.prototype.wipe = function () {
        return __awaiter(this, void 0, void 0, function () {
            var req;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0: return [4 /*yield*/, axios_1.default
                            .delete("".concat(this.options.url, "/clear"))
                            .then(function (r) {
                            return {
                                ok: true,
                            };
                        })
                            .catch(function (r) {
                            return {
                                ok: false,
                            };
                        })];
                    case 1:
                        req = _a.sent();
                        if (!req.ok) {
                            (0, Logger_1.default)("Error clearing/wiping store. Please check the db server terminal for more information.").error();
                        }
                        else {
                            (0, Logger_1.default)("Elix store has been cleared/wiped.").info();
                        }
                        return [2 /*return*/];
                }
            });
        });
    };
    Client.prototype.remove = function (key) {
        return __awaiter(this, void 0, void 0, function () {
            var req;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0: return [4 /*yield*/, axios_1.default
                            .delete("".concat(this.options.url, "/delete?key=").concat(key))
                            .then(function (r) {
                            return {
                                ok: true,
                                data: r.data,
                            };
                        })
                            .catch(function (r) {
                            return {
                                ok: false,
                                reason: r,
                            };
                        })];
                    case 1:
                        req = _a.sent();
                        if (!req.ok) {
                            (0, Logger_1.default)("Error deleted pair in store. Please check the db server terminal for more information.").error();
                        }
                        else {
                            (0, Logger_1.default)("Pair: ".concat(key, " has been deleted.")).info();
                        }
                        return [2 /*return*/];
                }
            });
        });
    };
    Client.prototype.seed = function () {
        var _this = this;
        this.options.seedData.forEach(function (seed) {
            _this.set(seed.key, seed.value);
        });
    };
    return Client;
}());
exports.default = Client;
