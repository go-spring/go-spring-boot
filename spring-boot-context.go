/*
 * Copyright 2012-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package SpringBoot

import (
	"github.com/go-spring/spring-core"
)

// 思来想去很多遍，看着这个空空的 ApplicationContext 越发感觉多余，
// 但是我也想不清楚未来到底需不需要更多功能，幸好 GoLang 提供了类型
// 等价定义的语法，如果未来没有新的用法被探索出来，就去掉下面的实现。
type ApplicationContext = SpringCore.SpringContext

// ApplicationContext Application 上下文，这个接口现在只是继承
// 了 SpringContext 并没有添加新的功能，为什么不去掉它呢？第一个肯
// 定历史原因，另外不能排除未来可能添加新功能，然后从领域模型角度考虑
// 使用 ApplicationContext 对应 ApplicationEvent 也更为合理。
//type ApplicationContext interface {
//	SpringCore.SpringContext
//}
//
// defaultApplicationContext ApplicationContext 的默认实现
//type defaultApplicationContext struct {
//	// 导出 ApplicationContext 接口
//	_ ApplicationContext `export:""`
//
//	// 导出 SpringCore.SpringContext 接口
//	SpringCore.SpringContext `export:""`
//}
