# language_used_by_so
A generic service which gathers questions asked on StackOverflow by language tag
and presents regressional data, trends and other useful stats on a configurable
timeseries collection point.

## stat_consumer
This app is responsible simply for gathering the data and batching it up into
kafka for processing.

## kafka (& zookeeper)
To act as a stream platform for scalability. Obviously the traffic of this application
is small, but it is merely for demonstrative purposes.

## redis
Used by cAdvisor and by the language apps for caching and sequencing.

## grafana
For performance, stat and general visualisations.

## prometheus
A metric store for observability.

## cadvisor
Monitor container performances, ultimately exposed in grafana in addition to the cAdvisor UI.