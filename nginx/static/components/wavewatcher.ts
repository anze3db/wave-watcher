import {Component, bootstrap, View, For, If} from "angular2/angular2";

@Component({
    selector: 'wave-watcher'
})
@View({
    templateUrl: "components/wavewatcher.template",
    directives: [For, If]
})
export class WaveWatcher {
    loading: boolean;

    constructor() {
      this.loading = true;
      // Emulate loading data
      setTimeout(function(){
        this.loading = false;
      }.bind(this), 1500)
    }
}
