import {Component, bootstrap, View, For, If} from "angular2/angular2";
import {Counter} from "components/counter"
import {Forecasts} from "models/forecats";


@Component({
    selector: 'wave-watcher'
})
@View({
    templateUrl: "components/wavewatcher.template",
    directives: [For, If, Counter]
})
export class WaveWatcher {
    loading: boolean;
    forecasts;

    constructor() {
      this.loading = true;
      // Emulate loading data
      Forecasts.findAll().then(function(d) {
        this.loading = false;
        this.forecasts = d;
      }.bind(this));
    }
}
