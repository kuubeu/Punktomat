let header = document.querySelector("header");
let filters = document.querySelector(".filters");
let navheight = header.clientHeight + filters.clientHeight;

// let slidermin = document.querySelector(".slider-min");
// let slidermax = document.querySelector(".slider-max");

// let min = 20, max = 200;

document.documentElement.style.setProperty('--nav-size', navheight + 'px');

// slidermin.oninput = s => s.target.value > slidermax.value ? slidermax.value = s.target.value : min = s.target.value;