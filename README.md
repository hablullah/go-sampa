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

> **Disclaimer**: I'm not an astronomer or physicist. I'm just an amateur with some interests toward celestial bodies.
>
> **Notice**: tests in this package do not represent calculation accuracy. It just a baseline so the future change in this package can be tracked to make sure it will give similar results with the current calculation.

This package has been compared to several other apps i.e. [Time and Date][timedate], [Sunrise Sunset Calendars][ssc] and [Accurate Times][accut]. I've found that the results are accurate to within a minute. So, while the results of this package is in seconds, don't expect it to be exactly accurate to seconds, and instead it should be treated as rounding suggestions.

Between the three comparison apps that I mentioned before, I mostly use [Accurate Times][accut] as baseline since it's used in official capacity (to some extend) by Jordan and Indonesia. You can check the comparison result between this package and Accurate Times in spreadsheets in [`doc`](doc/) directory.

For most days, the differences is only in seconds. However, for some days you will see that there are days where the difference suddenly spiked compared to previous days. This can be seen for Sunrise in Tromso at 17th May 2022 and Moon transit in Jakarta at 24th February 2022.

I'm not sure what is the reason, but one of the possibility is because the difference in algorithm between this package and Accurate Times (which uses VSOP87). Fortunately, our results are correct when compared with [Time and Date][timedate], so hopefully this package should be good enough to use.

## Additional Notes

While calculating rise and set time, this package will uses your elevation as one of the calculation parameter. If it's very high, it can affect the times by a couple of minutes thanks to atmospheric refraction. Thanks to this, most apps that I know decided to set elevation to zero, which means every locations will be treated as located in sea level.

Since we are talking about rise and set time, apparently there are no algorithms that could exactly calculate when the Sun will [rise][when-rise] and [set][when-set], and I assume it's the same for the Moon. With that said, please note that the calculated rise and set time might not be accurate to the reality.

Finally, this package doesn't account to Daylight Saving Team (which personally I believe is useless and should be discarded), so adjust your timezone accordingly.

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
[when-rise]: https://skyandtelescope.org/astronomy-news/we-dont-really-know-when-the-sun-rises/
[when-set]: https://aty.sdsu.edu/explain/sunset_time.html
[ssc]: https://www.sunrisesunset.com/
[mit]: http://choosealicense.com/licenses/mit/
