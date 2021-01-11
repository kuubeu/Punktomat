<template
  ><div class="page-container">
    <md-app md-mode="fixed">
      <md-app-toolbar class="md-primary">
        <span class="md-title">Punktomat</span>
      </md-app-toolbar>
      <md-app-content>
        <md-table v-model="info" md-card>
          <md-table-toolbar>
            <h1 class="md-title">Lista czasopism naukowych</h1>
          </md-table-toolbar>

          <md-table-row slot="md-table-row" slot-scope="{ item }">
            <md-table-cell md-label="ID" md-sort-by="ID">{{
              item.ID
            }}</md-table-cell>
            <md-table-cell md-label="Punkty" md-sort-by="points">{{
              item.points
            }}</md-table-cell>
            <md-table-cell md-label="Tytuł" md-sort-by="title">{{
              item.title
            }}</md-table-cell>
            <md-table-cell md-label="ISSN" md-sort-by="issn">{{
              item.issn
            }}</md-table-cell>
            <md-table-cell md-label="Kategorie">
              <div class="tags">
                <div
                  class="chip"
                  v-bind:key="item.issn + category"
                  v-for="category in item.Categories"
                >
                  {{ category }}
                </div>
              </div>
            </md-table-cell>
          </md-table-row>
        </md-table>
      </md-app-content>
    </md-app>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "Main",
  data() {
    return {
      info: [
        {
          ID: 1,
          title: "2D Materials",
          issn: "2053-1583",
          "e-issn": "2053-1583",
          secondTitle: "2D Materials",
          secondIssn: "",
          "secondE-issn": "2053-1583",
          points: 140,
          Categories: [
            "architektura i urbanistyka",
            "automatyka, elektronika i elektrotechnika",
            "inżynieria biomedyczna",
            "inżynieria lądowa i transport",
            "inżynieria materiałowa",
            "inżynieria mechaniczna",
            "nauki farmaceutyczne",
            "nauki chemiczne",
            "nauki fizyczne",
          ],
        },
        {
          ID: 2,
          title: "3 Biotech",
          issn: "2190-572X",
          "e-issn": "2190-5738",
          secondTitle: "3 Biotech",
          secondIssn: "2190-572X",
          "secondE-issn": "2190-5738",
          points: 70,
          Categories: [
            "architektura i urbanistyka",
            "inżynieria biomedyczna",
            "inżynieria chemiczna",
            "inżynieria środowiska, górnictwo i energetyka",
            "nauki farmaceutyczne",
            "nauki o zdrowiu",
            "rolnictwo i ogrodnictwo",
            "zootechnika i rybactwo",
            "nauki biologiczne",
          ],
        },
        {
          ID: 3,
          title: "3C Empresa",
          issn: "2254-3376",
          "e-issn": "2254-3376",
          secondTitle: "",
          secondIssn: "",
          "secondE-issn": "",
          points: 20,
          Categories: [
            "architektura i urbanistyka",
            "nauki o zarządzaniu i jakości",
          ],
        },
      ],
    };
  },
  mounted() {
    axios.get("http://127.0.0.1:4000").then((response) => {
      this.info = response.data;
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
