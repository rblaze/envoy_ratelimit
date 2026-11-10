package stats

import (
	gostats "github.com/lyft/gostats"
	logger "github.com/sirupsen/logrus"

	"github.com/envoyproxy/ratelimit/src/settings"
	"github.com/envoyproxy/ratelimit/src/utils"
)

func NewStatManager(store gostats.Store, settings settings.Settings) *ManagerImpl {
	serviceScope := store.ScopeWithTags("ratelimit", settings.ExtraTags).Scope("service")
	return &ManagerImpl{
		store:                        store,
		rlStatsScope:                 serviceScope.Scope("rate_limit"),
		serviceStatsScope:            serviceScope,
		shouldRateLimitScope:         serviceScope.Scope("call.should_rate_limit"),
		checkRateLimitScope:          serviceScope.Scope("call.check_rate_limit"),
		consumeRateLimitScope:        serviceScope.Scope("call.consume_rate_limit"),
		sanitizeDescriptorMetricDots: settings.SanitizeDescriptorMetricDots,
	}
}

func (this *ManagerImpl) GetStatsStore() gostats.Store {
	return this.store
}

func (this *ManagerImpl) SanitizeDescriptorMetricDots() bool {
	return this.sanitizeDescriptorMetricDots
}

// Create new rate descriptor stats for a descriptor tuple.
// @param key supplies the fully resolved descriptor tuple.
// @return new stats.
func (this *ManagerImpl) NewStats(key string) RateLimitStats {
	ret := RateLimitStats{}
	logger.Debugf("Creating stats for key: '%s'", key)
	ret.Key = key
	key = utils.SanitizeStatName(key)
	ret.TotalHits = this.rlStatsScope.NewCounter(key + ".total_hits")
	ret.OverLimit = this.rlStatsScope.NewCounter(key + ".over_limit")
	ret.NearLimit = this.rlStatsScope.NewCounter(key + ".near_limit")
	ret.OverLimitWithLocalCache = this.rlStatsScope.NewCounter(key + ".over_limit_with_local_cache")
	ret.WithinLimit = this.rlStatsScope.NewCounter(key + ".within_limit")
	ret.ShadowMode = this.rlStatsScope.NewCounter(key + ".shadow_mode")
	ret.TotalNegativeHits = this.rlStatsScope.NewCounter(key + ".total_negative_hits")
	return ret
}

func (this *ManagerImpl) NewDomainStats(domain string) DomainStats {
	ret := DomainStats{}
	domain = utils.SanitizeStatName(domain)
	ret.NotFound = this.rlStatsScope.NewCounter(domain + ".domain_not_found")
	return ret
}

func newCallStats(scope gostats.Scope) ShouldRateLimitStats {
	ret := ShouldRateLimitStats{}
	ret.RedisError = scope.NewCounter("redis_error")
	ret.ServiceError = scope.NewCounter("service_error")
	ret.Inflight = scope.NewGauge("inflight")
	return ret
}

func (this *ManagerImpl) NewShouldRateLimitStats() ShouldRateLimitStats {
	return newCallStats(this.shouldRateLimitScope)
}

func (this *ManagerImpl) NewCheckRateLimitStats() ShouldRateLimitStats {
	return newCallStats(this.checkRateLimitScope)
}

func (this *ManagerImpl) NewConsumeRateLimitStats() ShouldRateLimitStats {
	return newCallStats(this.consumeRateLimitScope)
}

func (this *ManagerImpl) NewServiceStats() ServiceStats {
	ret := ServiceStats{}
	ret.ConfigLoadSuccess = this.serviceStatsScope.NewCounter("config_load_success")
	ret.ConfigLoadError = this.serviceStatsScope.NewCounter("config_load_error")
	ret.ShouldRateLimit = this.NewShouldRateLimitStats()
	ret.CheckRateLimit = this.NewCheckRateLimitStats()
	ret.ConsumeRateLimit = this.NewConsumeRateLimitStats()
	ret.GlobalShadowMode = this.serviceStatsScope.NewCounter("global_shadow_mode")
	return ret
}

func (this RateLimitStats) GetKey() string {
	return this.Key
}
