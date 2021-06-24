<template>
  <div id="app" class="container">
    <div>
      <div class="offset-md-3 py-5">
        <h1>Aplicación para consulta y carga de compradores, productos y transacciones</h1>
      </div>
      <b-card-group deck>
        <b-card title="Subir información" aling="left">
          <b-card-text>
            <form v-on:submit.prevent="insertBuyers">
              <div>
                <span>Subir información</span>
              </div>
              <div class="form-group">
                <button class="btn btn-primary">Agregar compradores del día</button>
              </div>
            </form>
            <form v-on:submit.prevent="insertProducts">
              <div class="form-group">
                <button class="btn btn-primary">Agregar productos del día</button>
              </div>
            </form>
          </b-card-text>
        </b-card>
        <b-card title="Generación de reportes" aling="center">
          <form v-on:submit.prevent="getBuyerReport">
            <div class="form-group">
              <span>Generación de reportes</span>
              <input v-model="BuyerID" type="text" id="website-input" placeholder="Enter buyer's ID" class="form-control">
            </div>
            <div class="form-group">
              <button class="btn btn-primary">Generar reporte de comprador</button>
            </div>
            <div class="form-group">
              <span>Reporte:</span>
              <pre class="report">{{ Report.BuyerReport }}</pre>
            </div>
          </form>
        </b-card>
        <b-card title="Información de todos los compradores:" aling="right">
          <form v-on:submit.prevent="getAllBuyers">
            <div>
              <span>Información de todos los compradores</span>
            </div>
            <div class="form-group">
              <button class="btn btn-primary">Cargar todos los compradores</button>
            </div>
            <div class="form-group">
              <span>Compradores:</span>
            </div>
            <div class="form-group">
              <pre>{{ Buyers }}</pre>
            </div>
          </form>
        </b-card>
      </b-card-group>
    </div>
  </div>
</template>

<script>

  import axios from 'axios';
  import Vue from 'vue'
  import { BootstrapVue } from 'bootstrap-vue'
  // Import Bootstrap an BootstrapVue CSS files (order is important)

  Vue.use(BootstrapVue)

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
  .subtitle1{
    margin-left: 55rem;
    margin-top: -20rem;
    margin-right: 2rem;
    font-size: 1.2rem;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
  }
  .subtitle2{
    margin-left: -4rem;
    margin-top: -20rem;
    margin-right: 2rem;
    margin-bottom: auto;
    font-size: 1.2rem;
  }
  .subtitle3{
    font-size: 1.2rem;
  }
  .report{
    margin-bottom: auto;
    margin-top: 0px;
  }
</style>
