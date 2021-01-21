<template>
  <div>
    <v-skeleton-loader
      v-if="loading"
      type="table-thead, table-tbody, table-tfoot"
    ></v-skeleton-loader>
    <v-data-table
      v-else
      v-model="selected"
      item-key="ID"
      :calculate-widths="false"
      :headers="headers"
      :items="magazines"
      :server-items-length="totalMagazines"
      :options.sync="options"
      :loading="loading"
      :items-per-page="20"
      :footer-props="{ 'items-per-page-options': [ 20, 50, 100 ] }"
      disable-filtering
      show-select
      fixed-header
      fixed-footer
      height="calc(100vh - 124px)"
      class="elevation-0 pa-0 ma-0"
    >
      <template v-slot:body="{ items }">
      <tbody>
        <tr
          v-for="item in items"
          :key="item.ID"
        >
          <td>
            <v-layout class="mt-n3 mb-n4">
              <v-checkbox
                v-model="selected"
                :value="item"
                off-icon="mdi-star-outline"
                on-icon="mdi-star"
                color="amber"
              >
              </v-checkbox>
            </v-layout>
          </td>
          <td>
            <v-tooltip 
              right
              nudge-left="8px"
              nudge-bottom="24px"
              :open-delay="200"
              transition="fade-transition"
            >
              <template v-slot:activator="{ on, attrs }">
                <v-btn
                  icon
                  v-bind="attrs"
                  v-on="on"
                  @click="searchForArticle(item.title, item.issn)"
                >
                  <v-icon>mdi-text-box-search-outline</v-icon>
                </v-btn>
              </template>
              <span>Search Google for this article</span>
            </v-tooltip>
          </td>
          <td class="pt-2 pb-2">
            {{ item.title }}
          </td>
          <td>
            <v-layout justify-center>
              <v-chip
                :color="getColor(item.points)"
                dark
              >
                {{ item.points }}
              </v-chip>
            </v-layout>
          </td>
          <td>
            {{ item.issn }}
          </td>
          <td>
            <v-chip-group
              max-width="100px"
              show-arrows
            >
              <v-chip
                v-for="chip in item.Categories"
                v-bind:key="chip.id"
                :ripple="false"
              >
                {{ chip }}
              </v-chip>
            </v-chip-group>
          </td>
        </tr>
      </tbody>
      </template>
    </v-data-table>
  </div>
</template>

<script>
  export default {
    name: 'Table',

    props: [ 'loading', 'magazines', 'totalMagazines' ],

    data () {
      return {
        headers: [
          { 
            text: 'Search',
            value: 'search',
            sortable: false,
            align: 'end',
            width: "70px",
          },
          {
            text: 'Title',
            value: 'name',
            width: "clamp(300px, 30vw, 600px)",
          },
          {
            text: 'Points',
            value: 'points',
            width: "88px",
          },
          { 
            text: 'ISSN',
            value: 'issn',
            sortable: false,
            width: "100px",
          },
          { 
            text: 'Categories',
            value: 'chips',
            sortable: false,
          },
        ],
        selected: [],
        options: {},
      }
    },

    watch: {
      options: {
        handler () {
          this.$emit('optionsChanged', this.options)
          // console.log(this.options)
        },
        deep: true,
      },
      selected: {
        handler () {
          this.$emit('selectionChanged', this.selected)
          // console.log(this.selected)
        }
      }
    },

    mounted() {
      // cache not transfered/implemented yet

      // if (JSON.parse(localStorage.getItem("selected"))) {
      //   this.selected = [...JSON.parse(localStorage.getItem("selected"))];
      //   this.selected.map((value) => {
      //     let i = 0;
      //     this.magazines.map((valueMagazine) => {
      //       if (valueMagazine.ID == value.ID) {
      //         this.$refs.selectableTable.selectRow(i);
      //       }
      //       i++;
      //     });
      //   });
      // }
      // this.$refs.selectableTable.selectRow(2);
    },
    
    methods: {
      getColor (item) {
        if (item > 100) return 'green'
        else if (item > 50) return 'orange'
        else return 'red'
      },
      searchForArticle (title, issn) {
        let url = `https://scholar.google.com/scholar?q=${encodeURIComponent(title)}+${encodeURIComponent(issn)}`
        window.open(url, '_blank')
      },
    },
  }
</script>
