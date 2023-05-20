# Go SAMPA [![Go Reference][doc-badge]][doc-url]

Go SAMPA is a package for calculating Sun and Moon position. SAMPA itself is acronym of **Sun and Moon Position Algorithm**. It uses algorithms from several sources:

- Algorithm for calculating Sun position is taken from [SPA][spa] paper by Ibrahim Reda and Afshin Andreas from NREL.
- Algorithm for calculating Moon position is taken from [SAMPA][sampa] paper which also created by Ibrahim Reda and Afshin Andreas.
- Algorithm for rise, set, transit time and phases of moon are taken from _Astronomical Algorithms_ by Jean Meeus.

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Features](#features)
- [API](#api)
- [Accuracy](#accuracy)
- [FAQ](#faq)
- [Resources](#resources)
- [License](#license)

## Features

- Calculate Sun and Moon position for a given time and position.
- Calculate rise, set and transit for Sun and Moon.
- Calculate time when Sun or Moon reach a certain elevation.
- Calculate moon phases around a specified date.
- Thanks to Go, it will seamlessly handle DST times.
- Fast and accurate enough for amateur use.

## API

You can check the Go documentations to see the available APIs. However, the main interests in this package are:

- Function `GetSunPosition` to calculate Sun location for the given time and location.
- Function `GetSunEvents` to calculate Sunrise, Sunset, transit and any other additional events for the given time and location.
- Function `GetMoonPosition` to calculate Moon location for the given time and location.
- Function `GetMoonEvents` to calculate Moonrise, Moonset, transit and any other additional events for the given time and location.
- Function `GetMoonPhases` to calculate the time for Moon phases around the specified date time.

Check out [`sample`](sample/) directory on how to use each function.

## Accuracy

> **Disclaimer**: I'm not an astronomer or physicist. I'm just an amateur with some interests toward celestial bodies.

This package has been extensively compared to [Time and Date][timedate], which uses algorithm based on work by the U.S. Naval Observatory and NASA's Jet Propulsion Laboratory. I've compared the calculations for following locations:

|       Name       |   Country    | Latitude | Longitude |      Timezone       |  Offset   | DST Offset |
| :--------------: | :----------: | :------: | :-------: | :-----------------: | :-------: | :--------: |
| Lord Howe Island |  Australia   | 31°33'S  | 159°05'E  | Australia/Lord_Howe | UTC+10:30 | UTC+11:00  |
|      Maputo      |  Mozambique  | 25°58'S  |  32°34'E  |    Africa/Maputo    | UTC+02:00 |            |
|    Amsterdam     | Netherlands  | 52°22'N  |  4°54'E   |         CET         | UTC+01:00 | UTC+02:00  |
|       Oslo       |    Norway    | 59°55'N  |  10°44'E  |         CET         | UTC+01:00 | UTC+02:00  |
|   Philipsburg    | Sint Maarten | 18°02'N  |  63°03'W  |        UTC-4        | UTC-04:00 |            |
|     New York     |     USA      | 40°43'N  |  74°01'W  |  America/New_York   | UTC-05:00 | UTC-04:00  |

The calculations can be seen in [`comparison`](scripts/comparison/) directory.

From the comparison, I've found that by average the results are accurate to within a minute for both Sun and Moon events. The difference are pretty small for area around equators, and become quite large for area with higher latitude (>45°) where the Sun might not rise or set for the entire day. However, I argue this package is still correct and accurate enough to use.

For example, the largest time difference for Sun event in our comparison data is for **astronomical dusk in Oslo at 22 April 2023**:

- TimeAndDate said the astronomical dusk will occur at "2023-04-22 01:00:16"
- Go SAMPA calculation result is "2023-04-22 01:01:04" (48 seconds difference)

However, if we compare the Sun elevation angle for both times using the official [SPA calculator][spa-calc], we'll get the following result:

- At "2023-04-22 01:00:16" (TimeAndDate result) the Sun elevation is -17.995° ([calc][spa-calc-tnd]).
- At "2023-04-22 01:01:04" (Go SAMPA result) the Sun elevation is -18.002° ([calc][spa-calc-go]).

Since astronomical dusk occured when Sun is 18° below horizons, both calculations are correct despite the huge time difference between them.

## FAQ

1. **Does the elevation affects calculation result?**

   Yes, it will affect the result for rise and set time. If the elevation is very high, it can affect the times by a couple of minutes thanks to atmospheric refraction. However, most apps that I know prefer to set elevation to zero, which means every locations will be treated as located in sea level.

2. **The rise and set times are different compared to reality!**

   Apparently there are no algorithms that could exactly calculate when the Sun will [rise][when-rise] and [set][when-set], and I assume it's the same for the Moon.

3. **Are the calculation results are accurate up to seconds?**

   While the results of this package are in seconds, as mentioned above there are several seconds different between this package and other famous calculators like TimeAndDate. So, it's better to not expect it to be exactly accurate to seconds, and instead treat it as minute rounding suggestions.

4. **Why are the sunrise and sunset times occured in different day?**

   In this package, the sunrise and sunset are connected to transit time (the time when Sun reach meridian). However, in area with higher latitude sometime the Sun will never rise nor set for the entire day. In this case, sunrise might occur yesterday and the sunset might occur tomorrow.

5. **Why are the moonrise and moonset times not calculated?**

   There are some days where Moon never reach the meridian. In those case the moonrise and moonset will not be calculated since in this package rise and set are chained to transit time.

## Resources

1. Reda, I.; Andreas, A. (2003). Solar Position Algorithm for Solar Radiation Applications. 55 pp.; NREL Report No. TP-560-34302, Revised January 2008. ([web][spa])
2. Reda, I. (2010). Solar Eclipse Monitoring for Solar Energy Applications Using the Solar and Moon Position Algorithms. 35 pp.; NREL Report No. TP-3B0-47681. ([web][sampa])
3. Meeus, J. (1998). Astronomical Algorithms, Second Edition.

## License

Go-Sampa is distributed using [MIT] license.

[doc-badge]: https://pkg.go.dev/badge/github.com/hablullah/go-sampa.svg
[doc-url]: https://pkg.go.dev/github.com/hablullah/go-sampa
[spa]: https://midcdmz.nrel.gov/spa/
[spa-calc]: https://midcdmz.nrel.gov/solpos/spa.html
[spa-calc-tnd]: https://midcdmz.nrel.gov/apps/spa.pl?syear=2023&smonth=4&sday=22&eyear=2023&emonth=4&eday=22&step=10&stepunit=1&otype=1&hr=1&min=0&sec=16&latitude=59.917&longitude=10.733&timezone=2&elev=0&press=1013.25&temp=10&dut1=0.0&deltat=64.797&azmrot=180&slope=0&refract=0.5667&field=40
[spa-calc-go]: https://midcdmz.nrel.gov/apps/spa.pl?syear=2023&smonth=4&sday=22&eyear=2023&emonth=4&eday=22&step=10&stepunit=1&otype=1&hr=1&min=1&sec=4&latitude=59.917&longitude=10.733&timezone=2&elev=0&press=1013.25&temp=10&dut1=0.0&deltat=64.797&azmrot=180&slope=0&refract=0.5667&field=40
[sampa]: https://midcdmz.nrel.gov/sampa/
[timedate]: https://www.timeanddate.com/
[when-rise]: https://skyandtelescope.org/astronomy-news/we-dont-really-know-when-the-sun-rises/
[when-set]: https://aty.sdsu.edu/explain/sunset_time.html
[mit]: http://choosealicense.com/licenses/mit/
