# Go SAMPA [![Go Reference][doc-badge]][doc-url]

Package for calculating Sun and Moon position using [SPA][spa] and [SAMPA][sampa] algorithm with several modifications. According to the papers, it should be usable between year -2000 to 6000.

## Features

- Calculate Sun and Moon position for a given time and position.
- Calculate rise, set and transit for Sun and Moon.
- Calculate time when Sun or Moon reach a certain elevation.
- Calculate moon phases around a specified date.
- Fast and (hopefully) accurate enough for amateur use.

## Algorithm

For Sun and Moon position, this library uses [SPA][spa] and [SAMPA][sampa] algorithm that developed by Ibrahim Reda and Afshin Andreas from NREL. For the rise, set and transit time (and also the other events) it uses algorithm from Jean Meeus.

## Accuracy

This package has been compared to several other apps e.g. [Accurate Times][accut], [Time and Date][timedate] and [Sunrise Sunset Calendars][ssc], and I've found that the results are accurate to within a minute. So, while the results of this package is in seconds, don't expect it to be exactly accurate to seconds, and instead it should be treated as rounding suggestions.

While calculating rise and set time, this package will use your elevation as one of the calculation parameter. If it's very high, it can affect the times by a couple of minutes thanks to atmospheric refraction.

Finally, this package doesn't account to Daylight Saving Team, so adjust your timezone accordingly.

## Resources

1. Reda, I.; Andreas, A. (2003). Solar Position Algorithm for Solar Radiation Applications. 55 pp.; NREL Report No. TP-560-34302, Revised January 2008. ([web][spa])
2. Reda, I. (2010). Solar Eclipse Monitoring for Solar Energy Applications Using the Solar and Moon Position Algorithms. 35 pp.; NREL Report No. TP-3B0-47681. ([web][sampa])
3. Meeus, J. (1998). Astronomical Algorithms, Second Edition.

## License

Go-Prayer is distributed using [MIT] license.

[doc-badge]: https://pkg.go.dev/badge/github.com/hablullah/go-sampa.svg
[doc-url]: https://pkg.go.dev/github.com/hablullah/go-sampa
[spa]: https://midcdmz.nrel.gov/spa/
[sampa]: https://midcdmz.nrel.gov/sampa/
[accut]: https://www.astronomycenter.net/accut.html?l=en
[timedate]: https://www.timeanddate.com/
[ssc]: https://www.sunrisesunset.com/
[mit]: http://choosealicense.com/licenses/mit/
