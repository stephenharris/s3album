var s3albumApiEndpoint = `/api`

var app = new Vue({
  el: "#viewer",
  data: {
    path: "pictures/",
    loading: true,
    selected: [],
    folders: [],
    objects: [],
  },
  mounted: async function () {
    this.fetchData()
  },

  methods: {
    breadCrumbs: function () {
      parts = this.path.split("/");
      parts.pop();
      path = "";
      breadCrumbs = [];
      for (crumb of parts) {
        path = path + crumb + "/";
        breadCrumbs.push({
          name: crumb,
          path: path,
        });
      }

      return breadCrumbs;
    },
    downloadSelected: async function () {
      this.loading = true;
      UIkit.modal("#downloading-photos").show();
      await bulkDownload(this.selected)
      this.loading = false;
      UIkit.modal("#downloading-photos").hide();
    },
    clearSelection: function () {
      this.selected = [];
    },
    confirmDelete: async function (event) {
      this.loading = true;
      await deleteObjects(this.selected);
      this.clearSelection();
      await this.fetchData();
      UIkit.modal("#delete-photos").hide();
    },
    fetchData: async function () {
      this.loading = true;
      var contents = await listContent(this.path);
      this.folders = contents.folders;
      this.objects = contents.objects;
      this.loading = false;
    },
    selectPath: function (event, path) {
      this.path = path;
      if (path !== this.spath) {
        this.objects = []
        this.folders = []
        this.fetchData();
      }
    },
    goUp: function (event, image) {
      var parts = this.path.replace(/\/$/, "").split("/");
      parts.pop();
      this.selectPath(event, parts.join("/") + "/");
    },
    reverseMessage: function () {
      this.message = this.message.split("").reverse().join("");
    },
  },
});

function listContent(path) {
  return fetch(`${s3albumApiEndpoint}/objects/${path}`)
  .then(response => response.json())
  .catch((err) => {
    alert("There was an error listing your albums: " + err.message);
  }); 
}

function deleteObjects(objects) {
  fetch(`${s3albumApiEndpoint}/objects/`, {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json; charset=UTF-8'
    },
    body: JSON.stringify(objects)
  })
  .catch((err) => {
    alert(err.msg)
    console.log(err);
  });
}

function bulkDownload(objects) {
  return fetch(`${s3albumApiEndpoint}/download/`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json; charset=UTF-8'
    },
    body: JSON.stringify(objects)
  })
  .catch((err) => {
    alert(err.msg)
    console.log(err);
  })
  .then( res => res.blob() )
  .then( blob => {
    var file = window.URL.createObjectURL(blob);
    window.location.assign(file);
  });
}