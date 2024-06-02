(function ($) {
  "use strict";

/*---------------------------
    Hero Slider One Activation
  -----------------------------------*/
  var swiper = new Swiper(".sliderone", {
    slidesPerView: 1,
    spaceBetween: 30,
    centeredSlides: true,
    loop:true,
    pagination: {
      el: ".swiper-pagination",
      clickable: true,
    },
    
    navigation: {
      nextEl: '.sliderone .swiper-button-next',
      prevEl: '.sliderone .swiper-button-prev'
  },
  });

  /*---------------------------
    Hero Slider Two Activation
  -----------------------------------*/
  var swiper = new Swiper(".slidertwo", {
    slidesPerView: 1,
    spaceBetween: 30,
    centeredSlides: true,
    pagination: {
      el: ".swiper-pagination",
      clickable: true,
    },
    
    navigation: false
  });

/* ----------------------------
    Tilt Animation 
-------------------------------*/
  $('.js-tilt').tilt({
      base: window,
      reset: true, 
      scale: 1.02, 
      reverse: true, 
      max: 15, 
      perspective: 3e3, 
      speed: 4e3
  });

/* ----------------------------
Odometer Activation 
-------------------------------*/
if( $('.odometer').length ){
  var elemOffset = $('.odometer').offset().top;
  var winHeight = $(window).height();
  if(elemOffset < winHeight){
    $('.odometer').each(function(){
      $(this).html($(this).data('count-to'));
    });
  }
  $(window).on('scroll', function(){
    var elemOffset = $('.odometer').offset().top;
    function winScrollPosition() {
      var scrollPos = $(window).scrollTop(),
        winHeight = $(window).height();
      var scrollPosition = Math.round(scrollPos + (winHeight / 1.2));
      return scrollPosition;
    }
    if ( elemOffset < winScrollPosition()) {
      $('.odometer').each(function(){
        $(this).html($(this).data('count-to'));
      });
    }	
  });
};

/*---------------------------
Testimonial Activation
-----------------------------------*/
var swiper = new Swiper(".testimonialone", {
slidesPerView: 1,
loop:true,
pagination: false,
allowTouchMove: true,
navigation: false,
pagination: {
  el: ".swiper-pagination",
  clickable: true,
}
});
/*---------------------------
Testimonial Two Activation
-----------------------------------*/
var swiper = new Swiper(".testimonialtwo",{
  loop:true,
  autoplay:false,
  slidesPerView:3,
  pagination:false,
  navigation: {
    nextEl: ".testimonial-slider-button-next",
    prevEl: ".testimonial-slider-button-prev",
  },
  breakpoints:{
    0: {
      slidesPerView: 1
    },
    480: {
      slidesPerView: 1
    },
    768: {
      slidesPerView: 2
    },
    992: {
      slidesPerView: 2
    },
    1200: {
      slidesPerView: 3
    }
  }
})
/*---------------------------
Vertical Testimonial Activation
-----------------------------------*/
var swiper = new Swiper(".vertical-testimonial", {
  loop:true,
  pagination: true,
  allowTouchMove: true,
  navigation: false,
  direction: "vertical",
  slidesPerView: 2,
  mousewheel: true,
  pagination: {
    el: ".swiper-pagination",
    clickable: true,
  },
  breakpoints:{
    0: {
      slidesPerView: 1
    },
    480: {
      slidesPerView: 2
    }
  }
});
/*---------------------------
  Brand Slider Activation
-----------------------------------*/
var swiper = new Swiper(".brand-slider",{
  loop:true,
  speed:800,
  autoplay:false,
  slidesPerView:5,
  spaceBetween:0,
  pagination:false,
  navigation:false,
  breakpoints:{
    0: {
      slidesPerView: 2,
      spaceBetween: 80,
    },
    480: {
      slidesPerView: 2,
      spaceBetween: 80,
    },
    768: {
      slidesPerView: 3,
      spaceBetween: 80,
    },
    992: {
      slidesPerView: 3,
      spaceBetween: 80,
    },
    1200: {
      slidesPerView: 5,
      spaceBetween: 80,
    }
  }
})
/*---------------------------
  Brand Slider Two Activation
-----------------------------------*/
var swiper = new Swiper(".brand-slider-two",{
  loop:true,
  speed:800,
  autoplay:false,
  slidesPerView:4,
  spaceBetween:0,
  pagination:false,
  navigation:false,
  breakpoints:{
    0: {
      slidesPerView: 2,
      spaceBetween: 20,
    },
    480: {
      slidesPerView: 2,
      spaceBetween: 60,
    },
    768: {
      slidesPerView: 3,
      spaceBetween: 80,
    },
    992: {
      slidesPerView: 3,
      spaceBetween: 100,
    },
    1200: {
      slidesPerView: 4,
      spaceBetween: 110,
    }
  }
})

/*---------------------------
  Work Carousel Activation
-----------------------------------*/
var swiper = new Swiper(".work-carousel",{
  loop:true,
  speed:800,
  autoplay:false,
  slidesPerView:4,
  spaceBetween:20,
  pagination:false,  
  navigation: {
    nextEl: '.work-carousel-wrapper .swiper-btn-next',
    prevEl: '.work-carousel-wrapper .swiper-btn-prev',
  },
  breakpoints:{
    0: {
      slidesPerView: 1
    },
    480: {
      slidesPerView: 2
    },
    768: {
      slidesPerView: 3
    },
    992: {
      slidesPerView: 3
    },
    1200: {
      slidesPerView: 4
    }
  }
})

/*-------------------------------
FancyBox Activation
-----------------------------------*/
$('.video-popup').fancybox();

  /*--
  Skillbar Instance
-----------------------------------*/

var deferJs = {
  i: function (e) {
      deferJs.d();
      deferJs.methods();
  },

  d: function (e) {
      this._window = $(window),
          this._document = $(document),
          this._body = $('body'),
          this._html = $('html')

  },

  methods: function (e) {
      deferJs.radialProgress();
  },

  radialProgress: function () {
      $('.radial-progress').waypoint(function () {
          $('.radial-progress').easyPieChart({
              lineWidth: 12,
              scaleLength: 0,
              rotate: 0,
              trackColor: false,
              lineCap: 'round',
              size: 178
          });
      }, {
          triggerOnce: true,
          offset: 'bottom-in-view'
      });
  }
}
deferJs.i();


})(jQuery);