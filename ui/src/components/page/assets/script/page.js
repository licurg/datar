export default {
  name: 'page',
  mounted: function() {
    this.GetUsers()
  },
  data () {
    return {
      users: null
    }
  },
  methods: {
    ChangeData: function (event) {
      if (event) {
        var rid = event.target.getAttribute("rid");
        var name = event.target.getAttribute("name");
        var value = event.target.value;
        if (rid != null && rid != "" && name != null && name!= "" && value != null && value != "") {
          if (name == "email") {
            if (!this.ValidateEmail(value)) {
              return false
            }
          }
          this.UpdateUser(rid, name, value)
        } else {
          this.GetUsers()
        }
      }
    },
    GetUsers: function () {
      this.$http.get('/api/getUsers').then(response => {
        this.users = response.body
      }, response => {
        console.log(response);
      });
    },
    PostUser: function (){
      this.$http.post('/api/postUser').then(response => {
        console.log("Користувач успішно доданий!");
        this.GetUsers()
      }, response => {
        console.log(response);
      });
    },
    DeleteUser: function (id){
      this.$http.delete('/api/deleteUser/' + id).then(response => {
        console.log("Користувача №" + id + " успішно видалено!");
        this.GetUsers()
      }, response => {
        console.log(response);
      });
    },
    UpdateUser: function (id, name, value){
      this.$http.put('/api/updateUser/' + id, {name: name, value: value}).then(response => {
        console.log("Данні успішно змінені!");
      }, response => {
        console.log(response);
      });
    },
    ValidateEmail: function (email) {
      var re = /[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,4}/igm;
      if (re.test(email)) {
        return true
      } else {
        return false
      }
    }
  }
}
