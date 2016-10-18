import { Component } from '@angular/core';

@Component({
  selector: 'my-app',
  template: `
    <h1>STATKUBE</h1>
    <h2>Engineer statistics</h2>
    <h3>{{devs.length}}</h3>
    <div>
        <label>from</label>
        <input type="date" value="{{start}}"/>
        <label>to</label>
        <input type="date" value="{{end}}"/>
    </div>
  `
})

export class AppComponent {
  devs = [new Developer("super dev", 2), new Developer("super dev2", 3)]
}

export class Developer {
  name: string;
  pr_count: number;
  constructor(name: string, pr_count: number) {
    this.name = name;
    this.pr_count = pr_count;
  }
}
