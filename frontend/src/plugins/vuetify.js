import Vue from "vue";
import Vuetify from "vuetify/lib/framework";
import pl from "vuetify/es5/locale/pl";
import colors from "vuetify/lib/util/colors";

Vue.use(Vuetify);

export default new Vuetify({
  lang: {
    locales: { pl },
    current: "pl",
  },
  theme: {
    themes: {
      light: {
        primary: colors.indigo,
      },
    },
  },
});
