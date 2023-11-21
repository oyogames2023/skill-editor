package storage

import (
	"changeme/pkg/constants"
	"changeme/pkg/preferences"
	"fmt"
	"gopkg.in/yaml.v3"
	"reflect"
	"strings"
	"sync"
)

type PreferencesStorage struct {
	storage *LocalStorage
	mu      sync.Mutex
}

func NewPreferencesStorage() *PreferencesStorage {
	return &PreferencesStorage{
		storage: NewLocalStorage("config.yaml"),
	}
}

func (p *PreferencesStorage) DefaultPreferencesStorage() preferences.DefaultPreferences {
	return preferences.NewDefaultPreferences()
}

func (p *PreferencesStorage) getPreferences() preferences.DefaultPreferences {
	data, err := p.storage.Load()
	if err != nil {
		return p.DefaultPreferencesStorage()
	}
	var result preferences.DefaultPreferences
	if err = yaml.Unmarshal(data, &result); err != nil {
		return p.DefaultPreferencesStorage()
	}
	return result
}

func (p *PreferencesStorage) GetPreferences() preferences.DefaultPreferences {
	p.mu.Lock()
	defer p.mu.Unlock()
	result := p.getPreferences()
	result.Behavior.AsideWith = max(result.Behavior.AsideWith, constants.DefaultAsideWith)
	result.Behavior.WindowWith = max(result.Behavior.WindowWith, constants.DefaultMinWindowWidth)
	result.Behavior.WindowHeight = max(result.Behavior.WindowHeight, constants.DefaultWindowHeight)
	return result
}

func (p *PreferencesStorage) setPreferences(pf *preferences.DefaultPreferences, key string, value any) error {
	parts := strings.Split(key, ".")
	if len(parts) > 0 {
		var reflectValue reflect.Value
		if reflect.TypeOf(pf).Kind() == reflect.Ptr {
			reflectValue = reflect.ValueOf(pf).Elem()
		} else {
			reflectValue = reflect.ValueOf(pf)
		}
		for i, part := range parts {
			part = strings.ToUpper(part[:1]) + part[1:]
			reflectValue = reflectValue.FieldByName(part)
			if reflectValue.IsValid() {
				if i == len(parts)-1 {
					reflectValue.Set(reflect.ValueOf(value))
					return nil
				}
			} else {
				break
			}
		}
	}
	return fmt.Errorf("invalid key path(%s)", key)
}

func (p *PreferencesStorage) savePreferences(pf *preferences.DefaultPreferences) error {
	data, err := yaml.Marshal(pf)
	if err != nil {
		return err
	}
	if err = p.storage.Store(data); err != nil {
		return err
	}
	return nil
}

func (p *PreferencesStorage) SetPreferences(pf *preferences.DefaultPreferences) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.savePreferences(pf)
}

func (p *PreferencesStorage) UpdatePreferences(values map[string]any) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	pf := p.getPreferences()
	for path, v := range values {
		if err := p.setPreferences(&pf, path, v); err != nil {
			return err
		}
	}
	return p.savePreferences(&pf)
}

func (p *PreferencesStorage) RestoreDefault() preferences.DefaultPreferences {
	p.mu.Lock()
	defer p.mu.Unlock()
	pf := p.DefaultPreferencesStorage()
	if err := p.savePreferences(&pf); err != nil {

	}
	return pf
}
