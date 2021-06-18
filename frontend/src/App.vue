<template>
  <div id="app" class="container">
    <div class="row">
      <div class="col-md-6 offset-md-3 py-5">
        <h1>Generate a thumbnail of a website</h1>

        <form v-on:submit.prevent="getBuyerReport">
          <div class="form-group">
            <input v-model="BuyerID" type="text" id="website-input" placeholder="Enter buyer's ID" class="form-control">
          </div>
          <div class="form-group">
            <button class="btn btn-primary">Generate!</button>
          </div>
          <div class="form-group">
            <span>Message: {{ Report }}</span>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>

 import axios from 'axios';

export default {
  name: 'App',
  
  data() { return {
    BuyerID: '',
    Report: '',
  } },

  methods: {
    getBuyerReport() {
      let part1 = "{BuyerReport(BuyerID:\\\""
      let part2 = "\\\"){Transactions, SameIp, Recomendations}}"
      let request = part1.concat(this.BuyerID, part2)
      console.log(request)
      axios.post("http://localhost:4000/grahpql", {
      headers: {
        'Content-Type' : 'text/plain'
      },
      data: {
        "query": request,
      },
      width: 1920,
      height: 1080,
      output: 'json',
      thumbnail_width: 300 
      })
      .then((response) => {
        this.Report = response.data.BuyerReport;
      })
      .catch((error) => {
        //console.log(response.data)
        window.alert(`The API returned an error: ${error}`);
      })
      console 
    } 
  }
}
</script>