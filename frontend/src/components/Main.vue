<template>
  <div class=".container-fluid">
    <b-navbar variant="faded" type="light">
      <b-navbar-brand tag="h1" class="mb-0">Punktomat</b-navbar-brand>
      <md-badge
        class="md-primary"
        v-bind:md-content="selected.length"
        md-dense
        v-if="selected"
      >
        <md-button @click="showModal" class="md-icon-button">
          <md-icon>article</md-icon>
        </md-button>
      </md-badge>
      <md-icon v-else>article</md-icon>
    </b-navbar>
    <div class="row justify-content-md-center">
      <b-skeleton-table
        v-if="loading"
        :rows="10"
        :columns="5"
        :table-props="{ striped: true }"
      ></b-skeleton-table>
      <Table
        v-else
        @clicked="getSelected()"
        v-bind:selectMode="selectMode"
        v-bind:loading="loading"
        v-bind:fields="fields"
        v-bind:magazines="magazines"
        v-bind:selected="selected"
      />
    </div>
    <b-modal
      ref="my-modal"
      size="xl"
      hide-footer
      title="Using Component Methods"
    >
      <div
        class="d-block text-center"
        v-for="select in selected"
        v-bind:key="select.issn"
      >
        <p ref="content">{{ select }}</p>
      </div>
      <b-button class="mt-3" variant="outline-danger" block @click="hideModal"
        >Zamknij</b-button
      >
      <b-button class="mt-2" variant="outline-warning" block @
         @click="genPDF">Generuj PDF</b-button
      >
    </b-modal>
  </div>
</template>

<script>
import axios from "axios";
import Table from "./Table";
import {jsPDF} from "jspdf";
//import html2canvas from "html2canvas"
//import pdfMake from "pdfmake";
//import {addHTML} from "html2canvas";
//import rasterizehtml from "rasterizehtml";

export default {
  name: "Main",
  components: { Table },
  data() {
    return {
      magazines: null,
      loading: true,
      fields: [
        "ID",
        "title",
        "issn",
        "points",
        { key: "Categories", label: "Kategorie" },
      ],
      selectMode: "multi",
      selected: [],
    };
  },
  methods: {
    getSelected: function() {
      this.selected = [...JSON.parse(localStorage.getItem("selected"))];
    },
    showModal() {
      this.$refs["my-modal"].show();
    },
    hideModal() {
      this.$refs["my-modal"].hide();
    },
    genPDF(){

 var articles_count=0;//amount of articles
var doc = new jsPDF();

this.selected.forEach(function(line, i){//renders firts page
 articles_count=i+1;
var splitTitle = doc.splitTextToSize(line.title, 120);
var splitSecondTitle = doc.splitTextToSize(line.secondTitle, 120);
    doc.text(15, 15 + (i * 32),
    
        
        "ID: " + line.ID +" "+
        "issn: "+ line.issn + " "+
        "points :"+line.points + " "+
        "second issn: "+line.secondIssn +"\n "
        
        );     

doc.line(15,17 + (i * 32),200,17 + (i * 32));
doc.text(15, 23 + (i * 32), "Title:");
doc.text(30, 23 + (i * 32), splitTitle);
doc.text(15, 41 + (i * 32), "Second title:");
doc.text(50, 41 + (i * 32), splitSecondTitle);

});


var pages_neded=(Math.round(articles_count/9)  +1  );


if(articles_count>9){//rendering second and else pages


for(var page=1;page<pages_neded;page++){

doc.addPage();
var sliced=this.selected.slice(9*page);
sliced.forEach(function(line, i){
var splitSecondTitle = doc.splitTextToSize(line.secondTitle, 140);
var splitTitle = doc.splitTextToSize(line.title, 180);

    doc.text(15, 15 + (i * 32),
    
        
        "ID: " + line.ID +" "+
        "issn: "+ line.issn + " "+
        "points :"+line.points + " "+
        "second issn: "+line.secondIssn +"\n "
        );     
doc.line(15,17 + (i * 32),200,17 + (i * 32));
doc.text(15, 28 + (i * 32), "Title:");
doc.text(30, 28 + (i * 32), splitTitle);   
doc.text(15, 41 + (i * 32), "Second title:");
doc.text(50, 41 + (i * 32), splitSecondTitle);
});
}
}


doc.save('ChosenArticles.pdf');
    }
  },
  mounted() {
    if (JSON.parse(localStorage.getItem("selected"))) {
      this.selected = [...JSON.parse(localStorage.getItem("selected"))];
    }
    axios.get("http://127.0.0.1:4000/api/scienceMagazine").then((response) => {
      this.magazines = response.data.results;
      this.total = response.data.total;
      this.loading = false;
    });
  },
};
</script>

<style>
.chip {
  -webkit-align-items: center;
  align-items: center;
  -webkit-border-radius: 15px;
  border-radius: 15px;
  -webkit-box-shadow: inset 0 0 0 1px rgb(100 121 143 / 30%);
  box-shadow: inset 0 0 0 1px rgb(100 121 143 / 30%);
  -webkit-box-sizing: border-box;
  box-sizing: border-box;
  color: #5f6368;
  display: -webkit-box;
  display: -webkit-flex;
  display: flex;
  height: 30px;
  margin-bottom: 4px;
  margin-right: 8px;
  /* max-width: 160px; */
  overflow: hidden;
  padding: 0 12px;
  white-space: nowrap;
}
.tags {
  display: -webkit-box;
  display: -webkit-flex;
  display: flex;
  min-height: 30px;
  padding-bottom: 2px;
  margin-top: 8px;
  max-width: calc(100vw - 184px);
  width: 100%;
  flex-wrap: wrap;
}
</style>
