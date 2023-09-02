/* 
	xBe by TEMPLATE STOCK
	templatestock.co @templatestock
	Released for free under the Creative Commons Attribution 3.0 license (templated.co/license)
*/

// jQuery for page scrolling feature - requires jQuery Easing plugin
$(function() {
    'use strict';

    // Back to top smooth scroll

    $('a.page-scroll').bind('click', function(event) {
        var $anchor = $(this);
        $('html, body').stop().animate({
            scrollTop: $($anchor.attr('href')).offset().top
        }, 1000, 'easeInOutExpo');
        event.preventDefault();
    });


    // Scroll to top animated button

    $(window).scroll(function() {
        if ($(this).scrollTop() > 100) {
            $('.scroll-up').fadeIn();
        } else {
            $('.scroll-up').fadeOut();
        }
    });


    // Highlight the top nav as scrolling occurs
    $('body').scrollspy({
        target: '.navbar-fixed-top'
    })

    // Closes the Responsive Menu on Menu Item Click
    $('.navbar-collapse ul li a').click(function() {
        $('.navbar-toggle:visible').click();
    });


    // Script for Portfolio Section

    $(document).ready(function() {

        // Slider top 


        $('#myCarousel-one').carousel({
            interval: 9000, //changes the speed
            keyboard: false,
        })


        //Team carousel

        $('#myCarousel-two').carousel({
            interval: 4000, //changes the speed
            keyboard: false,
        })

        //Clients carousel

        $('#myCarousel-three').carousel({
            interval: 4000, //changes the speed
            keyboard: false,
        })

        //Testimonials carousel

        $('#myCarousel-four').carousel({
            interval: 8000, //changes the speed
            keyboard: false,
        })

        // Google Maps

        google.maps.visualRefresh = true;

        var map;

        function initialize() {
            var geocoder = new google.maps.Geocoder();
            var address = $('#map-input').text(); /* change the map-input to your address */
            var mapOptions = {
                zoom: 15,
                mapTypeId: google.maps.MapTypeId.ROADMAP,
                scrollwheel: false
            };
            map = new google.maps.Map(document.getElementById('map-canvas'), mapOptions);

            if (geocoder) {
                geocoder.geocode({
                    'address': address
                }, function(results, status) {
                    if (status == google.maps.GeocoderStatus.OK) {
                        if (status != google.maps.GeocoderStatus.ZERO_RESULTS) {
                            map.setCenter(results[0].geometry.location);

                            var infowindow = new google.maps.InfoWindow({
                                content: address,
                                map: map,
                                position: results[0].geometry.location,
                            });

                            var marker = new google.maps.Marker({
                                position: results[0].geometry.location,
                                map: map,
                                title: address
                            });

                        } else {
                            alert("No results found");
                        }
                    }
                });
            }
        }
        google.maps.event.addDomListener(window, 'load', initialize);

        /* end google maps */


    });

    /* Animated Titles of Sections*/

    function isScrolledIntoView(elem) {
        var docViewTop = $(window).scrollTop();
        var docViewBottom = docViewTop + $(window).height();

        var elemTop = $(elem).offset().top;
        var elemBottom = elemTop + $(elem).height();

        return ((elemBottom <= docViewBottom) && (elemTop >= docViewTop));
    }

    $(window).scroll(function() {
        $('.section-heading, .section-subheading').each(function() {
            if (isScrolledIntoView(this) === true) {
                $(this).addClass('animated fadeInUp')
            }
        });


    });


    /*  Isotope Filter */

    (function() {
        var winDow = $(window);
        var $container = $('.portfolio-items');
        var $filter = $('.filter');

        try {
            $container.imagesLoaded(function() {
                $container.show();
                $container.isotope({
                    filter: '*',
                    layoutMode: 'masonry',
                    animationOptions: {
                        duration: 750,
                        easing: 'linear'
                    }
                });
            });
        } catch (err) {}

        winDow.bind('resize', function() {
            var selector = $filter.find('a.active').attr('data-filter');

            try {
                $container.isotope({
                    filter: selector,
                    animationOptions: {
                        duration: 750,
                        easing: 'linear',
                        queue: false,
                    }
                });
            } catch (err) {}
            return false;
        });

        $filter.find('a').click(function() {
            var selector = $(this).attr('data-filter');

            try {
                $container.isotope({
                    filter: selector,
                    animationOptions: {
                        duration: 750,
                        easing: 'linear',
                        queue: false,
                    }
                });
            } catch (err) {

            }
            return false;
        });


    }());


    /* Header shrink function*/

    $(function() {
        var shrinkHeader = 350;
        $(window).scroll(function() {
            var scroll = getCurrentScroll();
            if (scroll >= shrinkHeader) {
                $('.navbar').addClass('navbar-shrink');
            } else {
                $('.navbar').removeClass('navbar-shrink');
            }
        });

        function getCurrentScroll() {
            return window.pageYOffset || document.documentElement.scrollTop;
        }
    });

    /* Footer reveal function*/

    $('footer').footerReveal({
        zIndex: -100,
        shadow: true,
        shadowOpacity: 0.6

    });



});
