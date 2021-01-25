<template>
  <v-app>
    <v-app-bar
      class="elevation-0"
      color="white"
      max-height="64px"
      height="64px"
    >
      <v-toolbar-title v-if="!$vuetify.breakpoint.xs"
        >Punktomat</v-toolbar-title
      >

      <v-spacer v-if="!$vuetify.breakpoint.xs"></v-spacer>

      <v-toolbar dense max-width="720px" class="rounded-lg elevation-2">
        <v-tooltip bottom :open-delay="200" transition="fade-transition">
          <template v-slot:activator="{ on, attrs }">
            <v-btn
              icon
              v-bind="attrs"
              v-on="on"
              @click="searchBtnClicked()"
              class="mr-n3"
              style="z-index: 1"
            >
              <v-icon>mdi-magnify</v-icon>
            </v-btn>
          </template>
          <span>Wyszukaj</span>
        </v-tooltip>
        <v-text-field
          flat
          solo
          clearable
          single-line
          hide-details
          placeholder="Wyszukaj"
          v-model="searchText"
        ></v-text-field>
        <v-tooltip bottom :open-delay="200" transition="fade-transition">
          <template v-slot:activator="{ on, attrs }">
            <v-btn
              icon
              v-bind="attrs"
              v-on="on"
              @click.stop="filtersDialog = true"
              class="mr-n3"
            >
              <v-icon>mdi-filter-outline</v-icon>
            </v-btn>
          </template>
          <span>Pokaż filtry</span>
        </v-tooltip>
        <v-dialog
          v-model="filtersDialog"
          max-width="900"
          @click:outside="applyFilters"
        >
          <v-card>
            <v-card-title class="headline">
              Filtruj rezultaty
            </v-card-title>

            <v-subheader>Wybierz zakres punktowy</v-subheader>

            <v-card-text>
              <v-range-slider
                v-model="range"
                min="20"
                max="200"
                step="10"
                thumb-label="always"
                append-icon="mdi-thumb-up-outline"
                prepend-icon="mdi-thumb-down-outline"
                class="mt-8 mb-n8"
                :value="this.range"
              ></v-range-slider>
            </v-card-text>

            <v-subheader>Wybierz kategorie</v-subheader>

            <v-card-text>
              <v-chip-group v-model="filters" column multiple class="mb-n4">
                <v-chip
                  v-for="chip in allCategories"
                  v-bind:key="chip"
                  filter
                  outlined
                >
                  {{ chip }}
                </v-chip>
              </v-chip-group>
            </v-card-text>

            <v-card-actions>
              <v-btn text @click="clearFilters">Wyczyść</v-btn>

              <v-spacer></v-spacer>

              <v-btn
                text
                color="primary"
                @click="
                  filtersDialog = false;
                  applyFilters();
                "
                >Zastosuj</v-btn
              >
            </v-card-actions>
          </v-card>
        </v-dialog>
      </v-toolbar>

      <v-spacer v-if="!$vuetify.breakpoint.xs"></v-spacer>
      <v-dialog v-model="selectDialog" width="650" v-if="selected.length > 0">
        <template v-slot:activator="{ on, attrs }">
          <v-badge
            bottom
            offset-x="15"
            offset-y="20"
            color="green"
            v-bind:content="selected.length"
          >
            <v-btn icon v-bind="attrs" v-on="on" class="ml-2">
              <v-icon>mdi-file-download-outline</v-icon>
            </v-btn>
          </v-badge>
        </template>

        <v-card>
          <v-card-title class="headline grey lighten-2">
            Wybrane czasopisma
          </v-card-title>
          <MinTable v-bind:selectedMagazines="selected" />
          <v-divider></v-divider>
          <v-card-actions>
            <v-btn text @click="selectDialog = false">
              Zamknij
            </v-btn>
            <v-spacer></v-spacer>
            <v-btn
              color="primary"
              text
              @click="
                selectDialog = false;
                genPDF();
              "
            >
              Zapisz do PDF
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
      <v-tooltip v-else bottom :open-delay="200" transition="fade-transition">
        <template v-slot:activator="{ on, attrs }">
          <v-btn icon v-bind="attrs" v-on="on" class="ml-2">
            <v-icon>mdi-file-download-outline</v-icon>
          </v-btn>
        </template>
        <span>Zapisz ulubione do PDF</span>
      </v-tooltip>
      <v-tooltip bottom :open-delay="200" transition="fade-transition">
        <template v-slot:activator="{ on, attrs }">
          <v-btn icon v-bind="attrs" v-on="on">
            <v-icon>mdi-dots-vertical</v-icon>
          </v-btn>
        </template>
        <span>Więcej opcji</span>
      </v-tooltip>
    </v-app-bar>

    <v-main class="overflow-y-auto pa-0">
      <Table
        v-bind:loading="loading"
        v-bind:magazines="magazines"
        v-bind:totalMagazines="totalMagazines"
        @optionsChanged="tableOptionsChanged"
        @selectionChanged="tableSelectionChanged"
      />
    </v-main>
  </v-app>
</template>

<script>
import axios from "axios";
import Table from "./Table";
import MinTable from "./MinTable";
import { jsPDF } from "jspdf";

export default {
  name: "Main",

  components: { Table, MinTable },

  data() {
    return {
      selectDialog: false,
      time: null,
      loading: true,
      magazines: [],
      totalMagazines: 0,
      selected: [],
      range: [20, 200],
      filters: [],
      searchText: "",
      filtersDialog: false,
      firstLoad: true,
      // replace with api later
      allCategories: [
        "archeologia",
        "filozofia",
        "historia",
        "językoznawstwo",
        "literaturoznawstwo",
        "nauki o kulturze i religii",
        "nauki o sztuce",
        "architektura i urbanistyka",
        "automatyka, elektronika i elektrotechnika",
        "informatyka techniczna i telekomunikacja",
        "inżynieria biomedyczna",
        "inżynieria chemiczna",
        "inżynieria lądowa i transport",
        "inżynieria materiałowa",
        "inżynieria mechaniczna",
        "inżynieria środowiska, górnictwo i energetyka",
        "nauki farmaceutyczne",
        "nauki medyczne",
        "nauki o kulturze fizycznej",
        "nauki o zdrowiu",
        "nauki leśne",
        "rolnictwo i ogrodnictwo",
        "technologia żywności i żywienia",
        "weterynaria",
        "zootechnika i rybactwo",
        "ekonomia i finanse",
        "geografia społeczno-ekonomiczna i gospodarka przestrzenna",
        "nauki o bezpieczeństwie",
        "nauki o komunikacji społecznej i mediach",
        "nauki o polityce i administracji",
        "nauki o zarządzaniu i jakości",
        "nauki prawne",
        "nauki socjologiczne",
        "pedagogika",
        "prawo kanoniczne",
        "psychologia",
        "astronomia",
        "informatyka",
        "matematyka",
        "nauki biologiczne",
        "nauki chemiczne",
        "nauki fizyczne",
        "nauki o Ziemi i środowisku",
        "nauki teologiczne",
      ],
      options: {
        data: {
          categories: [],
          order: "title",
          orderDirection: "asc",
          minPoints: 0,
          maxPoints: 200,
          limit: 20,
          offset: 0,
          search: "",
        },
      },
    };
  },
  watch: {
    searchText: function() {
      if (this.time) {
        clearTimeout(this.time);
        this.time = null;
      }
      this.time = setTimeout(() => {
        this.searchBtnClicked();
      }, 700);
    },
    selected: function() {
      if (this.selected == 0) this.selectDialog = false;
    },
  },
  mounted() {
    if (localStorage.getItem("options") !== null) {
      this.options.data = JSON.parse(localStorage.getItem("options"));
      this.range[0] = this.options.data.minPoints;
      this.range[1] = this.options.data.maxPoints;
      if (this.options.data.categories.length != 0) {
        this.filters = this.options.data.categories.map((value) => {
          return this.allCategories.indexOf(value);
        });
      }
    }
    if (localStorage.getItem("search") !== null) {
      this.options.data.search = localStorage.getItem("search");
      this.searchText = localStorage.getItem("search");
    }
    this.getDataFromApi();
  },

  methods: {
    getDataFromApi() {
      this.loading = true;
      axios
        .post(
          `${process.env.VUE_APP_API_URL}/scienceMagazine`,
          this.options.data
        )
        .then((response) => {
          this.magazines = response.data.results;
          this.totalMagazines = response.data.total;
          this.loading = false;
        });
    },
    searchBtnClicked() {
      if (this.searchText == null) this.searchText = "";
      this.searchText = this.searchText.trim().replace(/\s+/g, " ");
      this.options.data.offset = 0;
      this.options.data.search = this.searchText;
      localStorage.setItem("search", this.searchText);
      this.getDataFromApi();
    },
    applyFilters() {
      this.options.data.minPoints = this.range[0];
      this.options.data.maxPoints = this.range[1];
      if (this.filters.length != 0)
        this.options.data.categories = this.filters.map((value) => {
          return this.allCategories[value];
        });
      else this.options.data.categories = [];
      localStorage.setItem("options", JSON.stringify(this.options.data));
      this.getDataFromApi();
    },
    clearFilters() {
      this.options.data.categories = [];
      this.options.data.minPoints = 20;
      this.options.data.maxPoints = 200;
      this.range[0] = 20;
      this.range[1] = 200;
      this.filters = [];
      this.getDataFromApi();
    },
    tableOptionsChanged(op) {
      const options = JSON.parse(JSON.stringify(op));
      const orderName = options.sortBy[0];
      this.options.data.order = orderName;
      if (options.sortDesc[0]) this.options.data.orderDirection = "desc";
      else this.options.data.orderDirection = "asc";
      const lim = options.itemsPerPage;
      this.options.data.limit = lim;
      this.options.data.offset = (options.page - 1) * lim;
      this.getDataFromApi();
    },
    tableSelectionChanged(selected) {
      this.selected = selected;
      localStorage.setItem("starred", JSON.stringify(this.selected));
    },
    genPDF() {
      var articles_count = 0; //amount of articles
      var doc = new jsPDF();
      this.selected.forEach(function(line, i) {
        //renders firts page
        articles_count = i + 1;
        var splitTitle = doc.splitTextToSize(line.title, 120);
        //var splitSecondTitle = doc.splitTextToSize(line.secondTitle, 120);
        doc.text(
          15,
          15 + i * 32,

          "ID: " +
            line.ID +
            " " +
            "issn: " +
            line.issn +
            " " +
            "points :" +
            line.points +
            " " +
            "second issn: " +
            line.secondIssn +
            "\n "
        );
        doc.line(15, 17 + i * 32, 200, 17 + i * 32);
        doc.text(15, 23 + i * 32, "Title:");
        doc.text(30, 23 + i * 32, splitTitle);
        doc.text(15, 41 + i * 32, "");
        doc.text(50, 41 + i * 32, "");
      });
      var pages_neded = Math.round(articles_count / 9) + 1;
      if (articles_count > 9) {
        //rendering second and else pages
        for (var page = 1; page < pages_neded; page++) {
          doc.addPage();
          var sliced = this.selected.slice(9 * page);
          sliced.forEach(function(line, i) {
            // var splitSecondTitle = doc.splitTextToSize(line.secondTitle, 140);
            var splitTitle = doc.splitTextToSize(line.title, 180);
            doc.text(
              15,
              15 + i * 32,

              "ID: " +
                line.ID +
                " " +
                "issn: " +
                line.issn +
                " " +
                "points :" +
                line.points +
                " " +
                "second issn: " +
                line.secondIssn +
                "\n "
            );
            doc.line(15, 17 + i * 32, 200, 17 + i * 32);
            doc.text(15, 28 + i * 32, "Title:");
            doc.text(30, 28 + i * 32, splitTitle);
            doc.text(15, 41 + i * 32, "");
            doc.text(50, 41 + i * 32, "");
          });
        }
      }
      doc.save("ChosenArticles.pdf");
    },
  },
};
</script>

<style lang="scss">
@import "~/src/sass/variables.scss";

html {
  overflow-y: hidden;
}
</style>
