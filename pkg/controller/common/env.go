/*
 * Copyright 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package common

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/runtime/inject"

	"github.com/gardener/gardener-extension-shoot-dns-service/pkg/controller/config"
)

type Env struct {
	name   string
	client client.Client
	ctx    context.Context
	config config.DNSServiceConfig
	logr.Logger
}

func NewEnv(name string, config config.DNSServiceConfig) *Env {
	return &Env{
		name:   name,
		ctx:    context.Background(),
		config: config,
	}
}

func (e *Env) Context() context.Context {
	return e.ctx
}

func (e *Env) Client() client.Client {
	return e.client
}

func (e *Env) Config() *config.DNSServiceConfig {
	return &e.config
}

// EntryLabelPrefix calculated the label prefix for dns entries managed for shoots of this garden
func (e *Env) EntryLabelPrefix() string {
	return fmt.Sprintf("%s.gardener.cloud/", e.config.GardenID)
}

func (e *Env) ShootId(namespace string) string {
	return fmt.Sprintf("%s%s", e.EntryLabelPrefix(), namespace)
}

// InjectFunc enables dependency injection into the actuator.
func (e *Env) InjectFunc(f inject.Func) error {
	return nil
}

// InjectClient injects the controller runtime client into the reconciler.
func (e *Env) InjectClient(client client.Client) error {
	e.client = client
	return nil
}

// InjectClient injects the controller runtime client into the reconciler.
func (e *Env) InjectLogger(l logr.Logger) error {
	e.Logger = l.WithName(e.name)
	return nil
}