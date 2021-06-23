<template>
  <div id="app" class="container">
    <div class="row">
      <div class="col-md-6 offset-md-3 py-5">
        <h1>Aplicación para consulta y carga de compradores, productos y transacciones</h1>

        <form v-on:submit.prevent="getBuyerReport">
          <div class="form-group">
            <span class="subtitle3">Generación de reportes</span>
            <input v-model="BuyerID" type="text" id="website-input" placeholder="Enter buyer's ID" class="form-control">
          </div>
          <div class="form-group">
            <button class="btn btn-primary">Generar reporte de comprador</button>
          </div>
          <div class="form-group">
            <span class="subtitle3">Reporte:</span>
            <pre>{{ Report.BuyerReport }}</pre>
          </div>
        </form>

      </div>
      <div class="subtitle">
        <form v-on:submit.prevent="getAllBuyers" class="subtitle">
          <div class="subtitle3">
            <span>Información de todos los compradores</span>
          </div>
          <div class="form-group  ">
            <button class="btn btn-primary">Cargar todos los compradores</button>
          </div>
          <div class="form-group">
            <span class="subtitle3">Compradores:</span>
          </div>
          <div class="form-group">
            <pre>{{ Buyers }}</pre>
          </div>
        </form>
      </div>

      <div class="subtitle2">
        <form v-on:submit.prevent="insertBuyers" class="subtitle3">
          <div class="subtitle">
            <span>Subir información</span>
          </div>
          <div class="form-group">
            <button class="btn btn-primary">Agregar compradores del día</button>
          </div>
        </form>
        <form v-on:submit.prevent="insertProducts" class="subtitle3">
          <div class="form-group">
            <button class="btn btn-primary">Agregar productos del día</button>
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
    Buyers: '',
  } },

  methods: {
    getBuyerReport() {
      // Captures the buyer's ID to send it into the API.
      let part1 = `{BuyerReport(BuyerID:\\"`
      let part2 = `\\"){Transactions, SameIp, Recomendations}}`
      let request = part1.concat(this.BuyerID, part2)
      console.log(request)
      axios({method: "POST",
             url: "http://localhost:4000/graphql",
             headers: {
              'Content-Type' : 'text/plain'
             },
             data: {
              "query": `{BuyerReport(BuyerID:"${this.BuyerID}"){Transactions, SameIp, Recomendations}}`
             }
      })
      .then((response) => {
        this.Report = response.data.data;
        console.log(response.data.data)
      })
      .catch((error) => {
        window.alert(`The API returned an error: ${error}`);
      })
    },
    
    getAllBuyers() {
      // Request all the buyers of the DB
      let part1 = "{Buyers{BuyerID, Name}}"
      console.log(part1)
      axios({method: "POST", 
             url: "http://localhost:4000/graphql",
             headers: {
                "Content-Type" : 'text/plain'
              },
             data: {"query": `{Buyers {BuyerID, Name}}`}
      })
      .then((response) => {
        this.Buyers = response.data.data;
      })
      .catch((error) => {
        window.alert(`The API returned an error: ${error}`);
      })
      console 
    },
    
    insertBuyers() {
      // Request to insert all present day's buyers
      let part1 = "{InsertBuyers{BuyerID}}"
      console.log(part1)
      axios({method: "POST",
             url: "http://localhost:4000/graphql",
             headers: {
               'Content-Type' : 'text/plain'
             },
             data: {
               "query": `{InsertBuyers{BuyerID}}`,
             },
      })
      .then((response) => {
        this.Report = response.data.Buyers;
        window.alert(`Compradores del día agregados correctamente.`);
      })
      .catch((error) => {
        window.alert(`The API returned an error: ${error}`);
      })
    }, 

    insertProducts() {
      // Request to insert all present day's products
      let part1 = "{InsertProducts{Product}}"
      console.log(part1)
      axios({method: "POST",
             url: "http://localhost:4000/graphql",
             headers: {
               'Content-Type' : 'text/plain'
             },
             data: {
               "query": `{InsertProducts {ProductID}}`,
             },
      })
      .then((response) => {
        this.Report = response.data.Buyers;
        window.alert(`Productos del día agregados correctamente.`);
      })
      .catch((error) => {
        window.alert(`The API returned an error: ${error}`);
      })
    } 
  }
}
</script>

<style>
  .subtitle{
    margin-left: 2rem;
    margin-top: 6rem;
    margin-right: -4rem;
    font-size: 1.2rem;
  }
  .subtitle2{
    margin-left: -8rem;
    margin-top: -23rem;
    margin-right: 2rem;
    font-size: 1.2rem;
  }
  .subtitle3{
    font-size: 1.2rem;
  }
</style>
