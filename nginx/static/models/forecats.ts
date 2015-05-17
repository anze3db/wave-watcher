function parseDates(data) {
  var date_fields = [
    "CreatedAt",
    "DeletedAt",
    "UpdatedAt",
    "LastUpdate",
    "NextUpdate",
    "Rise",
    "Set",
  ], key;

  for (key in data) {
    let val = data[key];
    if (date_fields.some(e => { return e === key})) {
      data[key] = new Date(val);
    } else if (Array.isArray(val)) {
      val.forEach(parseDates);
    } else if (typeof val === "object") {
      parseDates(val);
    }
  }
  return data;
}

export class Forecasts {

  static findAll() {
    return $.getJSON( "/api/updates").then(data => {
      return parseDates(data[0]);
    }).fail(err => {
      console.error("error loading /api/updates", err)
    });
  }
}
