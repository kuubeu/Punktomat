<template>
<v-app>
    <v-app-bar
      class="elevation-0"
      color="white"
      max-height="64px"
      height="64px"
    >
      <v-toolbar-title
        v-if="!$vuetify.breakpoint.xs"
      >Punktomat</v-toolbar-title>

      <v-spacer v-if="!$vuetify.breakpoint.xs"></v-spacer>

      <v-toolbar
        dense
        max-width="720px"
        class="rounded-lg elevation-2"
      >
        <v-tooltip
          bottom
          :open-delay="200"
          transition="fade-transition"
        >
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
          <span>Search</span>
        </v-tooltip>
        <v-text-field
          flat solo
          clearable
          single-line
          hide-details
          placeholder="Search"
          v-model="searchText"
        ></v-text-field>
        <v-tooltip
          bottom
          :open-delay="200"
          transition="fade-transition"
        >
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
          <span>Show filters</span>
        </v-tooltip>
        <v-dialog
          v-model="filtersDialog"
          max-width="900"
          @click:outside="applyFilters"
        >
          <v-card>
            <v-card-title class="headline">
              Filter results
            </v-card-title>

            <v-subheader>Select the range of points</v-subheader>

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
              ></v-range-slider>
            </v-card-text>

            <v-subheader>Choose categories</v-subheader>

            <v-card-text>
              <v-chip-group
                v-model="filters"
                column
                multiple
                class="mb-n4"
              >
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
              <v-spacer></v-spacer>

              <v-btn
                text
                @click="filtersDialog = false; applyFilters();"
              >Apply</v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>

      </v-toolbar>

      <v-spacer v-if="!$vuetify.breakpoint.xs"></v-spacer>

      <v-tooltip
        bottom
        :open-delay="200"
        transition="fade-transition"
      >
        <template v-slot:activator="{ on, attrs }">
          <v-btn
            icon
            v-bind="attrs"
            v-on="on"
            class="ml-2"
          >
            <!-- <v-icon>mdi-pdf-box</v-icon> -->
            <v-icon>mdi-file-download-outline</v-icon>
          </v-btn>
        </template>
        <span>Save selected articles as PDF</span>
      </v-tooltip>

      <v-tooltip
        bottom
        :open-delay="200"
        transition="fade-transition"
      >
        <template v-slot:activator="{ on, attrs }">
          <v-btn
            icon
            v-bind="attrs"
            v-on="on"
          >
            <v-icon>mdi-dots-vertical</v-icon>
          </v-btn>
        </template>
        <span>More options</span>
      </v-tooltip>

    </v-app-bar>

    <v-main
      class="overflow-y-auto pa-0"
    >
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
  import axios from 'axios';
  import Table from './Table';

  export default {
    name: 'Main',

    components: { Table },

    data () {
      return {
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
        allCategories: {
          101: "archeologia",
          102: "filozofia",
          103: "historia",
          104: "językoznawstwo",
          105: "literaturoznawstwo",
          106: "nauki o kulturze i religii",
          107: "nauki o sztuce",
          201: "architektura i urbanistyka",
          202: "automatyka, elektronika i elektrotechnika",
          203: "informatyka techniczna i telekomunikacja",
          204: "inżynieria biomedyczna",
          205: "inżynieria chemiczna",
          206: "inżynieria lądowa i transport",
          207: "inżynieria materiałowa",
          208: "inżynieria mechaniczna",
          209: "inżynieria środowiska, górnictwo i energetyka",
          301: "nauki farmaceutyczne",
          302: "nauki medyczne",
          303: "nauki o kulturze fizycznej",
          304: "nauki o zdrowiu",
          401: "nauki leśne",
          402: "rolnictwo i ogrodnictwo",
          403: "technologia żywności i żywienia",
          404: "weterynaria",
          405: "zootechnika i rybactwo",
          501: "ekonomia i finanse",
          502: "geografia społeczno-ekonomiczna i gospodarka przestrzenna",
          503: "nauki o bezpieczeństwie",
          504: "nauki o komunikacji społecznej i mediach",
          505: "nauki o polityce i administracji",
          506: "nauki o zarządzaniu i jakości",
          507: "nauki prawne",
          508: "nauki socjologiczne",
          509: "pedagogika",
          510: "prawo kanoniczne",
          511: "psychologia",
          601: "astronomia",
          602: "informatyka",
          603: "matematyka",
          604: "nauki biologiczne",
          605: "nauki chemiczne",
          606: "nauki fizyczne",
          607: "nauki o Ziemi i środowisku",
          701: "nauki teologiczne",
        },
        options: {
          data: {
            "categories": [],
            "order": "title",
            "orderDirection": "asc",
            "minPoints": 0,
            "maxPoints": 200,
            "limit": 20,
            "offset": 0,
            "search": ""
          }
        }
      }
    },

    mounted () {
      this.getDataFromApi()
      console.log(this.options.params)
    },

    methods: {
      getDataFromApi () {
        this.loading = true
        axios
          .post(`${process.env.VUE_APP_API_URL}/scienceMagazine`, this.options.data)
          .then(response => {
            this.magazines = response.data.results
            this.totalMagazines = response.data.total
            this.loading = false
          })
      },
      searchBtnClicked () {
        this.searchText = this.searchText.trim().replace(/\s+/g, ' ')
        // let q = this.searchText.trim().replace(/\s+/g, ' ')
        this.options.data.offset = 0
        this.options.data.search = this.searchText
        this.getDataFromApi()
      },
      applyFilters () {
        console.log(this.filters)
      },
      clearFilters () {
        this.options.data.categories = []
      },
      tableOptionsChanged (op) {
        const options = JSON.parse(JSON.stringify(op))
        const lim = options.itemsPerPage

        this.options.data.limit = lim
        this.options.data.offset = (options.page - 1) * lim
        this.getDataFromApi()
      },
      tableSelectionChanged (selected) {
        this.selected = selected
        localStorage.setItem('starred', JSON.stringify(this.selected))
      },
    },

  };
</script>

<style lang="scss">
  @import '~/src/sass/variables.scss';

  html {
    overflow-y: hidden;
  }
</style>

