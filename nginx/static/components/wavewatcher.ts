import {Component, bootstrap, View} from "angular2/angular2";

@Component({
    selector: 'wave-watcher'
})
@View({
    templateUrl: "components/wavewatcher.template"
})
export class WaveWatcher {
    name: string;

    constructor() {
      this.name = 'Alice';
    }
}
