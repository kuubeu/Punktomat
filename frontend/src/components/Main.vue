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
          color="purple darken-1"
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
              @click="filterBtnClicked()"
            >
              <v-icon>mdi-filter-outline</v-icon>
            </v-btn>
          </template>
          <span>Show filters</span>
        </v-tooltip>
        <v-dialog
          v-model="filtersDialog"
          max-width="720"
        >
          <v-card>
            <v-card-title class="headline">
              Filter magazines
            </v-card-title>

            <v-card-text>
              text 1
            </v-card-text>

            <v-card-actions>
              <v-spacer></v-spacer>

              <v-btn
                color="purple darken-1"
                text
                @click="filtersDialog = false"
              >Cancel</v-btn>

              <v-btn
                color="purple darken-1"
                text
                @click="filtersDialog = false"
              >Save</v-btn>
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
        searchText: "",
        filtersDialog: false,
        options: {
          data: {
            "categories": [
              ""
            ],
            "order": "title",
            "orderDirection": "desc",
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
      this.loadCached()
      this.getDataFromApi()
      console.log(this.options.params)
    },

    methods: {
      getDataFromApi () {
        this.loading = true
        axios
          .post(`${process.env.VUE_APP_API_URL}/scienceMagazine`, this.options)
          .then(response => {
            this.magazines = response.data.results
            this.totalMagazines = response.data.total
            this.loading = false
          })
      },
      loadCached () {
        if (JSON.parse(localStorage.getItem("selected"))) {
          this.selected = [...JSON.parse(localStorage.getItem("selected"))]
        }
      },
      searchBtnClicked () {
        this.searchText = this.searchText.trim().replace(/\s+/g, ' ')
        // let q = this.searchText.trim().replace(/\s+/g, ' ')
        if (this.searchText.length > 0)
          console.log(this.searchText)
      },
      saveFilters () {
        // filter
      },
      clearFilters () {
        // filter
      },
      tableOptionsChanged (options) {
        console.log(options)
        this.options = options
        // this.getDataFromApi()
      },
      tableSelectionChanged (selected) {
        this.selected = selected
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

