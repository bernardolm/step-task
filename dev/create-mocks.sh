#!/bin/bash

find . -name '*mock_*.go' -delete

# Exemple
# mockgen -destination=repository/testutil/mock_entity.go -package=mocks github.com/bernardolm/step-task/repository Entity
