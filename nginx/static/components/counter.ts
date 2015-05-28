import {Component, bootstrap, View, For, If} from "angular2/angular2";


@Component({
    selector: 'wv-counter',
    properties: {
      end: "end"
    }
})
@View({
    templateUrl: "components/counter.template",
    directives: [For, If]
})
export class Counter {
    loading: boolean;
    remaining: string;
    end: Date;
    constructor() {
      // TODO: Find a better way of doing this, this code should not be
      //       in the constructor at all
      setTimeout(this.update.bind(this), 0);
      setInterval(this.update.bind(this), 1000);
    }
    update() {
      console.log('Updating', moment(this.end).fromNow())
      this.remaining = moment(this.end).fromNow();
    }
}
