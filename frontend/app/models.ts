export class Developer {
  name: string;
  pr_count: number;
  constructor(name: string, pr_count: number) {
    this.name = name;
    this.pr_count = pr_count;
  }
};

export class TimeConstraint {
    start: string;
    end: string;
    name: string;
    constructor(start: string, end: string, name: string) {
        this.start = start;
        this.end = end;
        this.name = name;
    }
};
