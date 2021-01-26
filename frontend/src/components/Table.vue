<template>
  <div>
    <v-data-table
      v-model="selected"
      item-key="ID"
      :calculate-widths="false"
      :headers="headers"
      :items="magazines"
      :server-items-length="totalMagazines"
      :options.sync="options"
      :loading="loading"
      :items-per-page="20"
      :footer-props="{ 'items-per-page-options': [20, 50, 100] }"
      disable-filtering
      show-select
      fixed-header
      fixed-footer
      height="calc(100vh - 124px)"
      class="elevation-0 pa-0 ma-0 data-table"
    >
      <template v-slot:[`item.data-table-select`]="{ isSelected, select }">
        <v-layout class="mt-n3 mb-n4">
          <v-simple-checkbox
            off-icon="mdi-star-outline"
            on-icon="mdi-star"
            color="amber"
            :value="isSelected"
            @input="select($event)"
            :ripple="false"
          ></v-simple-checkbox>
        </v-layout>
      </template>
      <template v-slot:[`item.search`]="{ item }">
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
          <span>Wyszukaj w google</span>
        </v-tooltip>
      </template>
      <template v-slot:[`item.title`]="{ item }">
        <v-layout class="mt-2 mb-2">
          {{ item.title }}
        </v-layout>
      </template>
      <template v-slot:[`item.points`]="{ item }">
        <v-layout justify-center>
          <v-chip :color="getColor(item.points)" dark>
            {{ item.points }}
          </v-chip>
        </v-layout>
      </template>
      <template v-slot:[`item.issn`]="{ item }">
        {{ item.issn }}
      </template>
      <template v-slot:[`item.chips`]="{ item }">
        <v-chip-group
          v-if="item.Categories.length <= 2"
          class="mr-n2"
        >
          <v-chip
            v-for="chip in item.Categories"
            v-bind:key="chip.id"
            :ripple="false"
          >
            {{ chip }}
          </v-chip>
        </v-chip-group>

        <v-chip-group v-else class="mr-n2">
          <v-chip
            v-for="chip in item.Categories.slice(0, 2)"
            v-bind:key="chip.id"
            :ripple="false"
          >
            {{ chip }}
          </v-chip>
          <v-tooltip bottom :open-delay="200" transition="fade-transition">
            <template v-slot:activator="{ on, attrs }">
              <v-chip
                v-bind="attrs"
                v-on="on"
                :ripple="false"
              >
                <v-icon>mdi-dots-horizontal</v-icon>
              </v-chip>
            </template>
            <span>
              <p v-for="chip in item.Categories.slice(2)" v-bind:key="chip.id">
                {{ chip }}
              </p>
            </span>
          </v-tooltip>
        </v-chip-group>
      </template>
    </v-data-table>
  </div>
</template>

<script>
export default {
  name: "Table",

  props: ["loading", "magazines", "totalMagazines"],

  data() {
    return {
      headers: [
        {
          text: "Wyszukaj",
          value: "search",
          sortable: false,
          align: "center",
          width: "70px",
        },
        {
          text: "Tytuł",
          value: "title",
          width: "clamp(300px, 30vw, 600px)",
        },
        {
          text: "Punkty",
          value: "points",
          width: "88px",
        },
        {
          text: "ISSN",
          value: "issn",
          sortable: false,
          width: "100px",
        },
        {
          text: "Kategorie",
          value: "chips",
          sortable: false,
        },
      ],
      selected: [],
      options: {},
    };
  },

  watch: {
    options: {
      handler() {
        this.$emit("optionsChanged", this.options);
      },
      deep: true,
    },
    selected: {
      handler() {
        this.$emit("selectionChanged", this.selected);
      },
    },
  },

  mounted() {
    this.loadCachedData();
  },

  methods: {
    getColor(item) {
      if (item > 100) return "green";
      else if (item > 50) return "orange";
      else return "red";
    },
    searchForArticle(title, issn) {
      let url = `https://scholar.google.com/scholar?q=${encodeURIComponent(
        title
      )}+${encodeURIComponent(issn)}`;
      window.open(url, "_blank");
    },
    loadCachedData() {
      if (localStorage.getItem("starred"))
        this.selected = JSON.parse(localStorage.getItem("starred"));
    },
    scrollUp() {
      this.$vuetify.goTo(
        '.data-table tbody tr',
        {
          offset: 50,
          container: '.v-data-table__wrapper'
        }
      )
    }
  },
};
</script>

<style>
.v-application p {
  margin-bottom: 2px;
}
.v-application p:last-of-type {
  margin-bottom: 0;
}
/* small screen optimization */
@media only screen and (max-width: 480px) {
  .v-data-footer__select .v-select {
    margin: 13px 0 13px 18px !important;
  }
  .v-data-footer__pagination {
    margin: 0 24px 0 16px !important;
  }
  .v-data-footer {
    margin-right: 0 !important;
  }
}
@media only screen and (max-width: 380px) {
  .v-data-footer__select {
    font-size: 0;
  }
}
</style>
