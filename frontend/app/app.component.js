"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};
var core_1 = require('@angular/core');
var AppComponent = (function () {
    function AppComponent() {
        this.devs = [new Developer("super dev", 2), new Developer("super dev2", 3)];
    }
    AppComponent = __decorate([
        core_1.Component({
            selector: 'my-app',
            template: "\n    <h1>STATKUBE</h1>\n    <h2>Engineer statistics</h2>\n    <h3>{{devs.length}}</h3>\n    <div>\n        <label>from</label>\n        <input type=\"date\" value=\"{{start}}\"/>\n        <label>to</label>\n        <input type=\"date\" value=\"{{end}}\"/>\n    </div>\n  "
        }), 
        __metadata('design:paramtypes', [])
    ], AppComponent);
    return AppComponent;
}());
exports.AppComponent = AppComponent;
var Developer = (function () {
    function Developer(name, pr_count) {
        this.name = name;
        this.pr_count = pr_count;
    }
    return Developer;
}());
exports.Developer = Developer;
//# sourceMappingURL=app.component.js.map